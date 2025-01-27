import math

a, b = input().split(' ')

rate = float(a)
i = b[0]
b = int(b[1:])

sums = 0

if i == ">":
    for j in range(b+1):
        sums += math.e**(-1*rate) * (rate**j/math.factorial(j))
    print(round(1-sums, 4))
elif i == "<":
    for j in range(b):
        sums += math.e**(-1*rate) * (rate**j/math.factorial(j))
    print(round(sums, 4))
elif i == "=":
    print(round(math.e**(-1*rate) * (rate**b/math.factorial(b)), 4))

