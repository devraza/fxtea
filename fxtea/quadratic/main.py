import sys

a,b,c = list(map(int, input().split(' ')))

x1 = (-1*b - (b**2 - 4*a*c)**(1/2))/(2*a)
x2 = (-1*b + (b**2 - 4*a*c)**(1/2))/(2*a)

if x1 == x2:
    print(x1)
    sys.exit()
else:
    print(f"{x1}, {x2}")
