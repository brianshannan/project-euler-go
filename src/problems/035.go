package main

import (
	"fmt"

	"utils"
)

func is_prime_circular(n int, primes []int) bool {
	nl := 0
	a := n

	for a != 0 {
		nl++
		a /= 10
	}
	len_pow := utils.IntPow(10, nl-1)

	for i := 0; i < nl; i++ {
		if !utils.IsPrimeEffic(n, primes) {
			return false
		}
		// rotate
		n = 10*(n%len_pow) + n/len_pow
	}

	return true
}

func main() {
	max := 1000000
	count := 0
	primes := utils.GetPrimesList(max + 1)

	for i := 0; i < len(primes); i++ {
		// Could store so previously computed, like 79 and 97 don't each have
		// to go through function check, but it's fast enough
		if is_prime_circular(primes[i], primes) {
			count++
		}
	}

	fmt.Println(count)
}
