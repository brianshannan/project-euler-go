package main

import "fmt"

func main() {
	l := 1001
	n := (l - 1) / 2

	// 1 + sum(i(i-1)+1 + i*i+1 + i(i+1)+1 + i(i+2)+1, i=1, n)
	// 1 + sum(16*i*1 + 4*i + 4, i=1, n)
	// 1 + 16*n*(n+1)*(2n+1)/6 + 2*n*(n+1) + 4*n
	sum := 1 + (16*n*n*n+30*n*n+26*n)/3
	fmt.Println(sum)
}
