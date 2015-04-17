package utils

func GetPascalsTriangle(n int) [][]int {
    l := n + 1
    t := make([][]int, l)
    for i, _ := range t {
        t[i] = make([]int, l)
    }

    for i := 0; i < n; i++ {
        t[i][0] = 1
        t[i][i] = 1
    }

    for i := 2; i < l; i++ {
        for j := 1; j < i; j++ {
            t[i][j] = t[i - 1][j - 1] + t[i - 1][j]
        }
    }
    return t
}
