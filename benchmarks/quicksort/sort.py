import sys
sys.setrecursionlimit(200000)

N = 100000
arr = [0] * N
seed = 42
for i in range(N):
    seed = (seed * 1103515245 + 12345) % 2147483648
    arr[i] = seed % 1000000

def quicksort(lo, hi):
    if lo >= hi:
        return
    pivot = arr[hi]
    i = lo
    for j in range(lo, hi):
        if arr[j] <= pivot:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1
    arr[i], arr[hi] = arr[hi], arr[i]
    quicksort(lo, i - 1)
    quicksort(i + 1, hi)

quicksort(0, N - 1)

sorted_flag = 1
for i in range(1, N):
    if arr[i] < arr[i - 1]:
        sorted_flag = 0
print(sorted_flag)
print(arr[0])
print(arr[N - 1])
