package main

import (
	"fmt"
	"math/big"
)

func main() {
	// A bit shift probably would be faster but golang doesn't allow shifts that big
	b := big.NewInt(2)
	e := big.NewInt(1000)
	r := b.Exp(b, e, big.NewInt(0))

	s := r.String()
	sum := 0
	for _, value := range s {
		sum += int(value - '0')
	}
	fmt.Println(sum)
}
