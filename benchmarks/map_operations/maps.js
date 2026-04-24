const N = 100000;

const m = new Map();
for (let i = 0; i < N; i++) {
  m.set("key_" + i, i * 2);
}
console.log(m.size);

let found = 0;
for (let i = 0; i < N; i++) {
  if (m.has("key_" + i)) found++;
}
console.log(found);

let sum = 0;
for (const v of m.values()) {
  sum += v;
}
console.log(sum);

for (let i = 0; i < N; i += 2) {
  m.delete("key_" + i);
}
console.log(m.size);
