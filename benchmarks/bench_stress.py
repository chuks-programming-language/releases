# ============================================================
# PYTHON PERFORMANCE BENCHMARK SUITE
# All times in milliseconds
# ============================================================

import time

def fib(n):
    a, b = 0, 1
    for _ in range(n):
        a, b = b, a + b
    return a

def matmul_sim(size):
    total = 0
    for i in range(size):
        for j in range(size):
            for k in range(size):
                total = total + (i * k) - (j * k) + 1
    return total

def count_primes(limit):
    count = 0
    for n in range(2, limit):
        is_prime = True
        d = 2
        while d * d <= n:
            if n % d == 0:
                is_prime = False
            d += 1
        if is_prime:
            count += 1
    return count

def array_stress(n):
    arr = []
    for i in range(n):
        arr.append(i * i)
    s = 0
    for i in range(len(arr)):
        s += arr[i]
    return s

def map_stress(n):
    m = {}
    for i in range(n):
        m["key"] = i * i
    s = 0
    keys = list(m.keys())
    for i in range(len(keys)):
        s += m[keys[i]]
    return s

class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    def add(self, other):
        return Vector(self.x + other.x, self.y + other.y)
    def magnitude_squared(self):
        return self.x * self.x + self.y * self.y

def oop_stress(n):
    total = 0
    for i in range(n):
        v1 = Vector(i, i + 1)
        v2 = Vector(i + 2, i + 3)
        v3 = v1.add(v2)
        total += v3.magnitude_squared()
    return total

class Shape:
    def __init__(self, name):
        self.name = name
    def area(self):
        return 0

class Rectangle(Shape):
    def __init__(self, w, h):
        super().__init__("rect")
        self.w = w
        self.h = h
    def area(self):
        return self.w * self.h

def inheritance_stress(n):
    total = 0
    for i in range(n):
        r = Rectangle(i, i + 1)
        total += r.area()
    return total

def make_counter():
    count = [0]
    def increment():
        count[0] += 1
        return count[0]
    return increment

def closure_stress(n):
    counter = make_counter()
    result = 0
    for _ in range(n):
        result = counter()
    return result

def string_stress(n):
    s = ""
    for _ in range(n):
        s += "x"
    return len(s)

def fib_recursive(n):
    if n <= 1:
        return n
    return fib_recursive(n - 1) + fib_recursive(n - 2)

# ============ RUN ALL ============

def ms():
    return int(time.time() * 1000)

print("=== PYTHON PERFORMANCE BENCHMARK SUITE ===")

t0 = ms()
fib(1000000)
t1 = ms()
print("1_fib_iter_1M")
print(t1 - t0)

t2 = ms()
matmul_sim(100)
t3 = ms()
print("2_nested_loops_1M")
print(t3 - t2)

t4 = ms()
count_primes(50000)
t5 = ms()
print("3_primes_50K")
print(t5 - t4)

t6 = ms()
array_stress(100000)
t7 = ms()
print("4_array_100K")
print(t7 - t6)

t8 = ms()
map_stress(50000)
t9 = ms()
print("5_map_50K")
print(t9 - t8)

t10 = ms()
oop_stress(100000)
t11 = ms()
print("6_oop_100K")
print(t11 - t10)

t12 = ms()
inheritance_stress(100000)
t13 = ms()
print("7_inherit_100K")
print(t13 - t12)

t14 = ms()
closure_stress(1000000)
t15 = ms()
print("8_closure_1M")
print(t15 - t14)

t16 = ms()
string_stress(10000)
t17 = ms()
print("9_string_10K")
print(t17 - t16)

t18 = ms()
fib_recursive(30)
t19 = ms()
print("10_fib_rec_30")
print(t19 - t18)

print("TOTAL")
print(t19 - t0)
