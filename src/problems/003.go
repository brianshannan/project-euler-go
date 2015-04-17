package main

import (
	"fmt"
	"math"
)

func get_primes_list(max_num int) []int {
	nums := make([]bool, max_num)
	for i := 0; i < len(nums); i++ {
		nums[i] = true
	}

	num_primes := 0
	for i := 2; i < max_num; i++ {
		if nums[i] == true {
			num_primes++
			for j := i * i; j < max_num; j += i {
				nums[j] = false
			}
		}
	}

	primes := make([]int, max_num)
	count := 0
	for i := 2; i < len(nums); i++ {
		if nums[i] == true {
			primes[count] = i
			count++
		}
	}
	return primes[:count]
}

func is_prime(num int, primes []int) bool {
	for i := 0; i < len(primes) && primes[i] < num; i++ {
		if num%primes[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 600851475143
	sqrt_num := int(math.Sqrt(float64(num))) + 1

	// Getting the primes < sqrt_num is sufficient as
	// if a number n is composite, it must be divisible
	// by some number m < sqrt(n), so sqrt_num
	// is the upper bound
	primes := get_primes_list(sqrt_num)

	// Check factors greater than the square root
	for i := 2; i < sqrt_num; i++ {
		if num%i == 0 && is_prime(num/i, primes) {
			fmt.Println(num / i)
			return
		}
	}

	// Check factors less than the square root
	for i := sqrt_num; i > 0; i-- {
		if num%i == 0 && is_prime(i, primes) {
			fmt.Println(i)
			return
		}
	}
}
