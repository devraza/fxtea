digits = int(input())

total = 0

for i in range(1, digits+1):
    total += (1 + (1/i))**i

total = str(total).replace(".", "")
total = list(total)

total.insert(1, ".")

print(''.join(total))
