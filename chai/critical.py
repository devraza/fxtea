import math

def li_gamma(s, x):
    def integrand(t):
        return t**(s - 1) * math.exp(-t)
    n = 1000
    h = x / n
    integral = 0.5 * (integrand(0) + integrand(x))
    for i in range(1, n):
        integral += integrand(i * h)
    integral *= h
    return integral

def chi(df, alpha, tol=1e-6):
    low, high = 0, 100
    while high-low>tol:
        mid = (low+high)/2
        if li_gamma(df/2, mid/2) / math.gamma(df/2) < 1-alpha:
            low = mid
        else:
            high = mid
    return (low+high)/2

df, alpha = map(float, input().split(' '))

print(f"{chi(df,alpha):.4f}")
