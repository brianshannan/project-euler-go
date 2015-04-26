package main

import (
	"fmt"
	"utils"
)

func get_primes_consec(a int, b int, primes []int) int {
	count := 0
	for {
		r := count*count + a*count + b
		if !utils.IsPrimeEffic(r, primes) {
			return count
		}
		count++
	}
}

func main() {
	primes := utils.GetPrimesList(10000000)
	max_count := -1
	max_a := -1000
	max_b := -1000

	for a := -999; a < 1000; a++ {
		for b := 2; b < 1000; b++ {
			count := get_primes_consec(a, b, primes)
			if count > max_count {
				max_count = count
				max_a = a
				max_b = b
			}
		}
	}

	fmt.Println(max_a * max_b)
}
