import math

def gamma(x):
    if x == int(x):
        return math.factorial(int(x) - 1)
    p = [
        676.5203681218851, -1259.1392167224028, 771.32342877765313,
        -176.61502916214059, 12.507343278686905, -0.13857109526572012,
        9.9843695780195716e-6, 1.5056327351493116e-7
    ]
    g = 7
    if x < 0.5:
        return math.pi / (math.sin(math.pi * x) * gamma_function(1 - x))
    x -= 1
    a = 0.99999999999980993
    for i, val in enumerate(p):
        a += val / (x + i + 1)
    return math.sqrt(2 * math.pi) * (x + g + 0.5) ** (x + 0.5) * math.exp(-(x + g + 0.5)) * a

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
        if li_gamma(df/2, mid/2) / gamma(df/2) < 1-alpha:
            low = mid
        else:
            high = mid
    return (low+high)/2

df, alpha = map(float, input().split(' '))

print(f"{chi(df,alpha):.4f}")
