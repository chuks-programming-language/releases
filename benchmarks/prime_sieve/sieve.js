const N = 1000000;
const sieve = new Array(N).fill(true);
sieve[0] = false;
sieve[1] = false;

for (let i = 2; i * i < N; i++) {
  if (sieve[i]) {
    for (let j = i * i; j < N; j += i) {
      sieve[j] = false;
    }
  }
}

let count = 0;
for (let i = 0; i < N; i++) {
  if (sieve[i]) count++;
}
console.log(count);
