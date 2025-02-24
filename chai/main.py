n = int(input())

total = 0
n_total = 0

for _ in range(n):
    o = int(input())
    e = int(input())

    total += (o**2)/e

    n_total += e

print(round(total-n_total, 4))
