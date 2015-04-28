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
