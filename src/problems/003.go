package main

import (
	"fmt"
	"math"

	"utils"
)

func main() {
	num := 600851475143
	sqrt_num := int(math.Sqrt(float64(num))) + 1

	// Getting the primes < sqrt_num is sufficient as
	// if a number n is composite, it must be divisible
	// by some number m < sqrt(n), so sqrt_num
	// is the upper bound
	primes := utils.GetPrimesList(sqrt_num)

	// Check factors greater than the square root
	for i := 2; i < sqrt_num; i++ {
		if num%i == 0 && utils.IsPrime(num/i, primes) {
			fmt.Println(num / i)
			return
		}
	}

	// Check factors less than the square root
	for i := sqrt_num; i > 0; i-- {
		if num%i == 0 && utils.IsPrime(i, primes) {
			fmt.Println(i)
			return
		}
	}
}
