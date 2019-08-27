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
