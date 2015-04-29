package main

import (
	"fmt"

	"utils"
)

func right_truncate_prime(n int, primes []int) bool {
	for n != 0 {
		if !utils.IsPrimeEffic(n, primes) {
			return false
		}
		n /= 10
	}

	return true
}

func left_truncate_prime(n int, primes []int) bool {
	a := n
	nl := 0
	for a != 0 {
		nl++
		a /= 10
	}

	init_div := utils.IntPow(10, nl-1)
	for n != 0 {
		if !utils.IsPrimeEffic(n, primes) {
			return false
		}
		n %= init_div
		init_div /= 10
	}

	return true
}

func main() {
	// Guess on upper limit
	max := 1000000
	primes := utils.GetPrimesList(max)
	count := 0
	sum := 0

	for i := 4; i < len(primes); i++ {
		if right_truncate_prime(primes[i], primes) && left_truncate_prime(primes[i], primes) {
			count++
			sum += primes[i]
		}
        if count == 11 {
            break
        }
	}

	if count < 11 {
		fmt.Println("Didn't get to 11 possible, need a larger prime list")
	}
	fmt.Println(sum)
}
