package evaluator

import (
	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/object"
)

// Eval takes an ast.Node (starting with the RootNode) and traverses the AST.
// It switches on the node's type and recursively evaluates them appropriately
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// Statements
	case *ast.RootNode:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}
