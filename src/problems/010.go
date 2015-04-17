package main

import (
	"fmt"

	"utils"
)

func main() {
	sum := 0
	primes := utils.GetPrimesList(2000000)

	for _, value := range primes {
		sum += value
	}

	fmt.Println(sum)

}
