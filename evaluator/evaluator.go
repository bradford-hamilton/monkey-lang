package evaluator

import (
	"fmt"

	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/object"
)

// No need to create new true/false objects every time we encounter one, they will
// be the same. Let's reference them instead
var (
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
	Null  = &object.Null{}
)

// Eval takes an ast.Node (starting with the RootNode) and traverses the AST.
// It switches on the node's type and recursively evaluates them appropriately
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Statements
	case *ast.RootNode:
		return evalRootNode(node, env)

	case *ast.BlockStatement:
		return evalBlockStmt(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

	case *ast.ConstStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObj(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpr(node.Operator, right, node.Token.Line)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpr(node.Operator, left, right, node.Token.Line)

	case *ast.PostfixExpression:
		return evalPostfixExpr(env, node.Operator, node)

	case *ast.IfExpression:
		return evalIfExpr(node, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{
			Parameters: params,
			Body:       body,
			Env:        env,
		}

	case *ast.CallExpression:
		fn := Eval(node.Function, env)
		if isError(fn) {
			return fn
		}
		args := evalExprs(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(fn, args, node.Token.Line)

	case *ast.ArrayLiteral:
		elements := evalExprs(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpr(left, index, node.Token.Line)

	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	}

	return nil
}

func evalRootNode(rootNode *ast.RootNode, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range rootNode.Statements {
		result = Eval(stmt, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStmt(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range block.Statements {
		result = Eval(stmt, env)

		if result != nil {
			rt := result.Type()
			if rt == object.ReturnValueObj || rt == object.ErrorObj {
				return result
			}
		}
	}

	return result
}

func nativeBoolToBooleanObj(input bool) *object.Boolean {
	if input {
		return True
	}

	return False
}

// Coerce our different object types to booleans for truthy/falsey values
func coerceObjToNativeBool(o object.Object) bool {
	if rv, ok := o.(*object.ReturnValue); ok {
		o = rv.Value
	}

	switch obj := o.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.String:
		return obj.Value != ""
	case *object.Null:
		return false
	case *object.Integer:
		return obj.Value != 0
	case *object.Array:
		return len(obj.Elements) > 0
	case *object.Hash:
		return len(obj.Pairs) > 0
	default:
		return true
	}
}

func evalPrefixExpr(operator string, right object.Object, line int) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpr(right)
	case "-":
		return evalMinusPrefixOperatorExpr(right, line)
	default:
		return newError("Line %d: Unknown operator: %s%s", line, operator, right.Type())
	}
}

func evalBangOperatorExpr(right object.Object) object.Object {
	switch right {
	case True:
		return False
	case False:
		return True
	case Null:
		return True
	default:
		return False
	}
}

func evalMinusPrefixOperatorExpr(right object.Object, line int) object.Object {
	if right.Type() != object.IntegerObj {
		return newError("Line %d: Unknown operator: -%s", line, right.Type())
	}

	value := right.(*object.Integer).Value

	return &object.Integer{Value: -value}
}

func evalPostfixExpr(env *object.Environment, operator string, node *ast.PostfixExpression) object.Object {
	switch operator {
	case "++":
		val, ok := env.Get(node.Token.Literal)
		if !ok {
			return newError("Line %d: Token literal %s is unknown", node.Token.Line, node.Token.Literal)
		}

		arg, ok := val.(*object.Integer)
		if !ok {
			return newError("Line %d: Invalid left-hand side expression in postfix operation", node.Token.Line)
		}

		v := arg.Value
		env.Set(node.Token.Literal, &object.Integer{Value: v + 1})
		return arg

	case "--":
		val, ok := env.Get(node.Token.Literal)
		if !ok {
			return newError("Line %d: Token literal %s is unknown", node.Token.Line, node.Token.Literal)
		}

		arg, ok := val.(*object.Integer)
		if !ok {
			return newError("Line %d: Invalid left-hand side expression in postfix operation", node.Token.Line)
		}

		v := arg.Value
		env.Set(node.Token.Literal, &object.Integer{Value: v - 1})
		return arg

	default:
		return newError("Line %d: Unknown operator: %s", node.Token.Line, operator)
	}
}

func evalInfixExpr(operator string, left, right object.Object, line int) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpr(operator, left, right, line)
	case left.Type() == object.StringObj && right.Type() == object.StringObj:
		return evalStringInfixExpr(operator, left, right, line)
	case operator == "==":
		return nativeBoolToBooleanObj(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObj(left != right)
	case operator == "&&":
		return nativeBoolToBooleanObj(coerceObjToNativeBool(left) && coerceObjToNativeBool(right))
	case operator == "||":
		return nativeBoolToBooleanObj(coerceObjToNativeBool(left) || coerceObjToNativeBool(right))
	case left.Type() != right.Type():
		return newError("Line %d: Type mismatch: %s %s %s", line, left.Type(), operator, right.Type())
	default:
		fmt.Printf("%s %s %s", left.Type(), operator, right.Type())
		return newError("Line %d: Unknown operator: %s %s %s", line, left.Type(), operator, right.Type())
	}
}

func evalIfExpr(ifExpr *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ifExpr.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ifExpr.Consequence, env)
	} else if ifExpr.Alternative != nil {
		return Eval(ifExpr.Alternative, env)
	} else {
		return Null
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case Null:
		return false
	case True:
		return true
	case False:
		return false
	default:
		return true
	}
}

func evalIntegerInfixExpr(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToBooleanObj(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObj(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObj(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObj(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObj(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObj(leftVal != rightVal)
	default:
		return newError("Line %d: Unknown operator: %s %s %s", line, left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpr(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObj(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObj(leftVal != rightVal)
	default:
		return newError("Line %d: Unknown operator: %s %s %s", line, left.Type(), operator, right.Type())
	}

}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtinFn, ok := builtinFunctions[node.Value]; ok {
		return builtinFn
	}

	return newError("Line %d: Identifier not found: %s", node.Token.Line, node.Value)
}

func evalIndexExpr(left, index object.Object, line int) object.Object {
	switch {
	case left.Type() == object.ArrayObj && index.Type() == object.IntegerObj:
		return evalArrayIndexExpr(left, index)
	case left.Type() == object.HashObj:
		return evalHashIndexExpr(left, index, line)
	default:
		return newError("Line %d: Index operator not supported: %s", line, left.Type())
	}
}

func evalArrayIndexExpr(array, index object.Object) object.Object {
	arrayObj := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObj.Elements) - 1)

	if idx < 0 || idx > max {
		return Null
	}

	return arrayObj.Elements[idx]
}

func evalHashIndexExpr(hash, index object.Object, line int) object.Object {
	hashObj := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("Line %d: Unusable as a hash key: %s", line, index.Type())
	}

	pair, ok := hashObj.Pairs[key.HashKey()]
	if !ok {
		return Null
	}

	return pair.Value
}

func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("Line %d: Unusable as a hash key: %s", node.Token.Line, key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func evalExprs(exprs []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, expr := range exprs {
		evaluated := Eval(expr, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(function object.Object, args []object.Object, line int) object.Object {
	switch fn := function.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		if result := fn.Fn(args...); result != nil {
			return result
		}
		return Null
	default:
		return newError("Line %d: Not a function: %s", line, function.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		env.Set(param.Value, args[i])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func newError(msgWithFormatVerbs string, values ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(msgWithFormatVerbs, values...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ErrorObj
	}
	return false
}
