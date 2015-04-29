package main

import (
	"fmt"
	"math"

	"utils"
)

func has_pandigital_divisors(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 && utils.IsPandigital(i, n/i, n) {
			return true
		}
	}
	return false
}

func main() {
	// Let x, y, z be in N such that x * y = z
	// Let a, b, c be the number of digits in x, y, z respectively
	// As the combination of x, y, z is pandigital, a + b + c = 9
	// If c <= 3, a + b >= 6. The smallest such product of x, y satisying
	// this is 1 * 10000 -> c = 5 > 3. So c > 3 -> c >= 4.
	// If c >= 5, a + b <= 4. The largest such product of x, y satisfying
	// this is 99 * 99 -> c = 4 < 5. So c < 5 -> c <= 4.
	// As c <= 4 and c >= 4, c = 4 -> 1000 <= z <= 9999
	// (products need to be counted only once, so iterate through them)

	sum := 0
	for i := 1000; i <= 9999; i++ {
		if has_pandigital_divisors(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
