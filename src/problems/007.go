package main

import (
	"fmt"

	"utils"
)

func main() {
	primes := utils.GetPrimesList(1000000)
	fmt.Println(primes[10000])
}
