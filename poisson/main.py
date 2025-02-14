import math

a, b = input().split(' ')

rate = float(a)
i = b[0]

if 'CR' not in b:
    b = int(b[1:])
else:
    level = float(b[2:])
    cumulative = 0
    count = 0
    while round(cumulative, 4) != 1.:
        cumulative += math.e**(-1*rate) * (rate**count/math.factorial(count))
        print(f"{round(cumulative, 4)} {count}")
        count += 1

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

