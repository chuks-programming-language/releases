N = 1000000
sieve = [True] * N
sieve[0] = False
sieve[1] = False

i = 2
while i * i < N:
    if sieve[i]:
        j = i * i
        while j < N:
            sieve[j] = False
            j += i
    i += 1

count = sum(1 for x in sieve if x)
print(count)
