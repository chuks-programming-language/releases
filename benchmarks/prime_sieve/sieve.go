package main

import "fmt"

func main() {
    N := 1000000
    sieve := make([]bool, N)
    for i := range sieve {
        sieve[i] = true
    }
    sieve[0] = false
    sieve[1] = false

    for i := 2; i*i < N; i++ {
        if sieve[i] {
            for j := i * i; j < N; j += i {
                sieve[j] = false
            }
        }
    }

    count := 0
    for i := 0; i < N; i++ {
        if sieve[i] {
            count++
        }
    }
    fmt.Println(count)
}
