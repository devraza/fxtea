import math

a = list(map(float, input().split(' ')))
rate = a[0]

sums = 0

for i in range(1,len(a)):
    sums += round(math.e**(-1*rate) * (rate**a[i]/math.factorial(int(a[i]))), 4)

print(sums)
