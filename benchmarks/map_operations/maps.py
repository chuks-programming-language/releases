N = 100000

m = {}
for i in range(N):
    m["key_" + str(i)] = i * 2
print(len(m))

found = 0
for i in range(N):
    if "key_" + str(i) in m:
        found += 1
print(found)

total = sum(m.values())
print(total)

for i in range(0, N, 2):
    del m["key_" + str(i)]
print(len(m))
