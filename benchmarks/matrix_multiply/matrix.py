N = 200

A = [i % 100 for i in range(N * N)]
B = [(i * 3 + 7) % 100 for i in range(N * N)]
C = [0] * (N * N)

for i in range(N):
    for j in range(N):
        s = 0
        for k in range(N):
            s += A[i * N + k] * B[k * N + j]
        C[i * N + j] = s

checksum = sum(C[i * N + i] for i in range(N))
print(checksum)
