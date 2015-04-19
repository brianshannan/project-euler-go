package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// smallest thousand digit number
	max_num := "1" + strings.Repeat("0", 999)
	max := big.NewInt(0)
	max.SetString(max_num, 10)

	a := big.NewInt(1)
	b := big.NewInt(1)
	count := 2
	for b.Cmp(max) == -1 {
		next := big.NewInt(0)
		next.Add(a, b)
		a, b = b, next
		count++
	}

	fmt.Println(count)
}
