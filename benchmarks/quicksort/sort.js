const N = 100000;
const arr = new Array(N);
let seed = 42;
for (let i = 0; i < N; i++) {
  seed = (seed * 1103515245 + 12345) % 2147483648;
  arr[i] = seed % 1000000;
}

function quicksort(lo, hi) {
  if (lo >= hi) return;
  const pivot = arr[hi];
  let i = lo;
  for (let j = lo; j < hi; j++) {
    if (arr[j] <= pivot) {
      [arr[i], arr[j]] = [arr[j], arr[i]];
      i++;
    }
  }
  [arr[i], arr[hi]] = [arr[hi], arr[i]];
  quicksort(lo, i - 1);
  quicksort(i + 1, hi);
}

quicksort(0, N - 1);

let sorted = 1;
for (let i = 1; i < N; i++) {
  if (arr[i] < arr[i - 1]) sorted = 0;
}
console.log(sorted);
console.log(arr[0]);
console.log(arr[N - 1]);
