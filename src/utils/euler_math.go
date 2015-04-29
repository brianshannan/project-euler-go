package utils

func Gcd(a int, b int) int {
    if a == b {
        return a
    } else if a > b {
        return Gcd(b, a - b)
    } else {
        return Gcd(a, b - a)
    }
}

func Factorial(n int) int {
    prod := 1
    for i := 2; i <= n; i++ {
        prod *= i
    }
    return prod
}

func IntPow(a int, x int) int {
    if x == 0 {
        return 1
    } else if x == 1 {
        return a
    }

    if x % 2 == 0 {
        return IntPow(a*a, x/2)
    }
    return a * IntPow(a*a, (x-1)/2)
}
