const N = 200;

const A = new Array(N * N);
const B = new Array(N * N);
const C = new Array(N * N).fill(0);

for (let i = 0; i < N * N; i++) {
  A[i] = i % 100;
  B[i] = (i * 3 + 7) % 100;
}

for (let i = 0; i < N; i++) {
  for (let j = 0; j < N; j++) {
    let sum = 0;
    for (let k = 0; k < N; k++) {
      sum += A[i * N + k] * B[k * N + j];
    }
    C[i * N + j] = sum;
  }
}

let checksum = 0;
for (let i = 0; i < N; i++) {
  checksum += C[i * N + i];
}
console.log(checksum);
