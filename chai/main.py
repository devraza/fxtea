n = int(input())

total = 0

for _ in range(n):
    o = float(input())
    e = float(input())

    total += ((o - e)**2)/e

print(round(total, 4))
