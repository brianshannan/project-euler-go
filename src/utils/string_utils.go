package utils

func Reverse(s string) string {
    runes := []rune(s)
    l := len(runes)
    for i := 0; i < l/2; i++ {
        runes[i], runes[l - i - 1] = runes[l - i - 1], runes[i]
    }
    return string(runes)
}
