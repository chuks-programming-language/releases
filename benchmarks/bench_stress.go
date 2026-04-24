package main

import (
	"fmt"
	"strings"
	"time"
)

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func matmulSim(size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				sum = sum + (i*k) - (j*k) + 1
			}
		}
	}
	return sum
}

func countPrimes(limit int) int {
	count := 0
	for n := 2; n < limit; n++ {
		isPrime := true
		for d := 2; d*d <= n; d++ {
			if n%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	return count
}

func arrayStress(n int) int {
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, i*i)
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func mapStress(n int) int {
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		m["key"] = i * i
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

type Vector struct {
	x, y int
}

func NewVector(x, y int) *Vector { return &Vector{x, y} }

func (v *Vector) Add(other *Vector) *Vector {
	return NewVector(v.x+other.x, v.y+other.y)
}

func (v *Vector) MagnitudeSquared() int {
	return v.x*v.x + v.y*v.y
}

func oopStress(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		v1 := NewVector(i, i+1)
		v2 := NewVector(i+2, i+3)
		v3 := v1.Add(v2)
		total += v3.MagnitudeSquared()
	}
	return total
}

type Shape struct {
	name string
}

func NewShape(name string) *Shape { return &Shape{name} }
func (s *Shape) Area() int       { return 0 }

type Rectangle struct {
	*Shape
	w, h int
}

func NewRectangle(w, h int) *Rectangle {
	return &Rectangle{Shape: NewShape("rect"), w: w, h: h}
}

func (r *Rectangle) Area() int {
	return r.w * r.h
}

func inheritanceStress(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		r := NewRectangle(i, i+1)
		total += r.Area()
	}
	return total
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func closureStress(n int) int {
	counter := makeCounter()
	result := 0
	for i := 0; i < n; i++ {
		result = counter()
	}
	return result
}

func stringStress(n int) int {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("x")
	}
	return sb.Len()
}

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func ms(d time.Duration) int64 { return d.Milliseconds() }

func main() {
	fmt.Println("=== GO PERFORMANCE BENCHMARK SUITE ===")

	t0 := time.Now()
	fib(1000000)
	t1 := time.Now()
	fmt.Println("1_fib_iter_1M")
	fmt.Println(ms(t1.Sub(t0)))

	t2 := time.Now()
	matmulSim(100)
	t3 := time.Now()
	fmt.Println("2_nested_loops_1M")
	fmt.Println(ms(t3.Sub(t2)))

	t4 := time.Now()
	countPrimes(50000)
	t5 := time.Now()
	fmt.Println("3_primes_50K")
	fmt.Println(ms(t5.Sub(t4)))

	t6 := time.Now()
	arrayStress(100000)
	t7 := time.Now()
	fmt.Println("4_array_100K")
	fmt.Println(ms(t7.Sub(t6)))

	t8 := time.Now()
	mapStress(50000)
	t9 := time.Now()
	fmt.Println("5_map_50K")
	fmt.Println(ms(t9.Sub(t8)))

	t10 := time.Now()
	oopStress(100000)
	t11 := time.Now()
	fmt.Println("6_oop_100K")
	fmt.Println(ms(t11.Sub(t10)))

	t12 := time.Now()
	inheritanceStress(100000)
	t13 := time.Now()
	fmt.Println("7_inherit_100K")
	fmt.Println(ms(t13.Sub(t12)))

	t14 := time.Now()
	closureStress(1000000)
	t15 := time.Now()
	fmt.Println("8_closure_1M")
	fmt.Println(ms(t15.Sub(t14)))

	t16 := time.Now()
	stringStress(10000)
	t17 := time.Now()
	fmt.Println("9_string_10K")
	fmt.Println(ms(t17.Sub(t16)))

	t18 := time.Now()
	fibRecursive(30)
	t19 := time.Now()
	fmt.Println("10_fib_rec_30")
	fmt.Println(ms(t19.Sub(t18)))

	fmt.Println("TOTAL")
	fmt.Println(ms(t19.Sub(t0)))
}
