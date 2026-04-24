package main

import "fmt"

var arr []int

func quicksort(lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	quicksort(lo, i-1)
	quicksort(i+1, hi)
}

func main() {
	N := 100000
	arr = make([]int, N)
	seed := 42
	for i := 0; i < N; i++ {
		seed = (seed*1103515245 + 12345) % 2147483648
		arr[i] = seed % 1000000
	}

	quicksort(0, N-1)

	sorted := 1
	for i := 1; i < N; i++ {
		if arr[i] < arr[i-1] {
			sorted = 0
		}
	}
	fmt.Println(sorted)
	fmt.Println(arr[0])
	fmt.Println(arr[N-1])
}
