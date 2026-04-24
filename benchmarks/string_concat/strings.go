package main

import (
	"fmt"
	"strings"
)

func main() {
	N := 100000

	var b strings.Builder
	for i := 0; i < N; i++ {
		b.WriteByte('a')
	}
	fmt.Println(len(b.String()))

	var b2 strings.Builder
	for i := 0; i < 10000; i++ {
		b2.WriteString("abc")
	}
	fmt.Println(len(b2.String()))
}
