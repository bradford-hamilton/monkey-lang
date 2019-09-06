let globalNum = 10;

let sum = func(a, b) {
  let c = a + b;
  c + globalNum;
};

let outer = func() {
  sum(1, 2) + sum(3, 4) + globalNum;
};

print(outer() + globalNum);

if (true && false) {
  print("shouldn't happen");
} else {
  print("should happen");
}

// Single line comment
print("Nice") // Or like this

const cantChangeMe = "neat";
print(cantChangeMe);

const moduloResult = (21 % 8) % (8 % 3);
print(moduloResult);

const fullName = func(firstName, lastName) {
  return firstName + " " + lastName;
}

print(fullName("Thurman", "Murman"));

let anotherGlobalNum = 55;
anotherGlobalNum++;
print(anotherGlobalNum);

let postfixPlusFunc = func() {
  let localNum = 99;
  localNum++;
  return localNum;
}

print(postfixPlusFunc());

let postfixMinusFunc = func() {
  let localNum = 99;
  localNum--;
  return localNum;
}

print(postfixMinusFunc());

/*
  ignored multiline comment print("will not print")
*/

print(5 >= 5)
print(5 <= 5)
print(5 > 5)
print(5 < 5)

/* ignored multiline comment */

print("string comparison below should print true true true")
const one = "one"
const two = "two"
print("monkey" == "monkey")
print("monkey" != "lang")
print(one != two)