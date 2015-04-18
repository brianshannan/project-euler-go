package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(100)
	for i := 99; i > 0; i-- {
		n = n.Mul(n, big.NewInt(int64(i)))
	}

	sum := big.NewInt(0)
	d := big.NewInt(0)
	for n.Cmp(big.NewInt(0)) == 1 {
		d.Mod(n, big.NewInt(10))
		sum.Add(sum, d)
		n.Div(n, big.NewInt(10))
	}

	fmt.Println(sum)
}
