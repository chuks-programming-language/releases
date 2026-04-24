// ============================================================
// NODE.JS PERFORMANCE BENCHMARK SUITE
// All times in milliseconds
// ============================================================

function fib(n) {
  let a = 0, b = 1;
  for (let i = 0; i < n; i++) {
    const tmp = a + b;
    a = b;
    b = tmp;
  }
  return a;
}

function matmul_sim(size) {
  let sum = 0;
  for (let i = 0; i < size; i++) {
    for (let j = 0; j < size; j++) {
      for (let k = 0; k < size; k++) {
        sum = sum + (i * k) - (j * k) + 1;
      }
    }
  }
  return sum;
}

function count_primes(limit) {
  let count = 0;
  for (let n = 2; n < limit; n++) {
    let is_prime = true;
    for (let d = 2; d * d <= n; d++) {
      if (n % d === 0) {
        is_prime = false;
      }
    }
    if (is_prime) count++;
  }
  return count;
}

function array_stress(n) {
  const arr = [];
  for (let i = 0; i < n; i++) {
    arr.push(i * i);
  }
  let sum = 0;
  for (let i = 0; i < arr.length; i++) {
    sum += arr[i];
  }
  return sum;
}

function map_stress(n) {
  const m = {};
  for (let i = 0; i < n; i++) {
    m["key"] = i * i;
  }
  let sum = 0;
  const keys = Object.keys(m);
  for (let i = 0; i < keys.length; i++) {
    sum += m[keys[i]];
  }
  return sum;
}

class Vector {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
  add(other) {
    return new Vector(this.x + other.x, this.y + other.y);
  }
  magnitudeSquared() {
    return this.x * this.x + this.y * this.y;
  }
}

function oop_stress(n) {
  let total = 0;
  for (let i = 0; i < n; i++) {
    const v1 = new Vector(i, i + 1);
    const v2 = new Vector(i + 2, i + 3);
    const v3 = v1.add(v2);
    total += v3.magnitudeSquared();
  }
  return total;
}

class Shape {
  constructor(name) { this.name = name; }
  area() { return 0; }
}

class Rectangle extends Shape {
  constructor(w, h) {
    super("rect");
    this.w = w;
    this.h = h;
  }
  area() { return this.w * this.h; }
}

function inheritance_stress(n) {
  let total = 0;
  for (let i = 0; i < n; i++) {
    const r = new Rectangle(i, i + 1);
    total += r.area();
  }
  return total;
}

function make_counter() {
  let count = 0;
  return function() { return ++count; };
}

function closure_stress(n) {
  const counter = make_counter();
  let result = 0;
  for (let i = 0; i < n; i++) {
    result = counter();
  }
  return result;
}

function string_stress(n) {
  let s = "";
  for (let i = 0; i < n; i++) {
    s += "x";
  }
  return s.length;
}

function fib_recursive(n) {
  if (n <= 1) return n;
  return fib_recursive(n - 1) + fib_recursive(n - 2);
}

// ============ RUN ALL ============

console.log("=== NODE.JS PERFORMANCE BENCHMARK SUITE ===");

let t0 = Date.now();
fib(1000000);
let t1 = Date.now();
console.log("1_fib_iter_1M");
console.log(t1 - t0);

let t2 = Date.now();
matmul_sim(100);
let t3 = Date.now();
console.log("2_nested_loops_1M");
console.log(t3 - t2);

let t4 = Date.now();
count_primes(50000);
let t5 = Date.now();
console.log("3_primes_50K");
console.log(t5 - t4);

let t6 = Date.now();
array_stress(100000);
let t7 = Date.now();
console.log("4_array_100K");
console.log(t7 - t6);

let t8 = Date.now();
map_stress(50000);
let t9 = Date.now();
console.log("5_map_50K");
console.log(t9 - t8);

let t10 = Date.now();
oop_stress(100000);
let t11 = Date.now();
console.log("6_oop_100K");
console.log(t11 - t10);

let t12 = Date.now();
inheritance_stress(100000);
let t13 = Date.now();
console.log("7_inherit_100K");
console.log(t13 - t12);

let t14 = Date.now();
closure_stress(1000000);
let t15 = Date.now();
console.log("8_closure_1M");
console.log(t15 - t14);

let t16 = Date.now();
string_stress(10000);
let t17 = Date.now();
console.log("9_string_10K");
console.log(t17 - t16);

let t18 = Date.now();
fib_recursive(30);
let t19 = Date.now();
console.log("10_fib_rec_30");
console.log(t19 - t18);

console.log("TOTAL");
console.log(t19 - t0);
