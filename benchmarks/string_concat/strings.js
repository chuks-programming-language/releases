const N = 100000;

let result = "";
for (let i = 0; i < N; i++) {
  result += "a";
}
console.log(result.length);

let pattern = "";
for (let i = 0; i < 10000; i++) {
  pattern += "abc";
}
console.log(pattern.length);
