package main

import (
	"fmt"
	"strconv"
)

func main() {
	N := 100000

	m := make(map[string]int)
	for i := 0; i < N; i++ {
		m["key_"+strconv.Itoa(i)] = i * 2
	}
	fmt.Println(len(m))

	found := 0
	for i := 0; i < N; i++ {
		if _, ok := m["key_"+strconv.Itoa(i)]; ok {
			found++
		}
	}
	fmt.Println(found)

	sum := 0
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)

	for i := 0; i < N; i += 2 {
		delete(m, "key_"+strconv.Itoa(i))
	}
	fmt.Println(len(m))
}
