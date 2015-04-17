package main

import (
	"fmt"

	"utils"
)

func factorial(n int) int {
	fact := 1
	for i := 2; i < n; i++ {
		fact *= i
	}
	return fact
}

func main() {
	n := 20
	fmt.Println(utils.GetPascalsTriangle(2 * n)[2*n][n])
}
