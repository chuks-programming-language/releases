package main

import "fmt"

func main() {
N := 200

A := make([]int, N*N)
B := make([]int, N*N)
C := make([]int, N*N)

for i := 0; i < N*N; i++ {
A[i] = i % 100
B[i] = (i*3 + 7) % 100
}

for i := 0; i < N; i++ {
for j := 0; j < N; j++ {
sum := 0
for k := 0; k < N; k++ {
sum += A[i*N+k] * B[k*N+j]
}
C[i*N+j] = sum
}
}

checksum := 0
for i := 0; i < N; i++ {
checksum += C[i*N+i]
}
fmt.Println(checksum)
}
