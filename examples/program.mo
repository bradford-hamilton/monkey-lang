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