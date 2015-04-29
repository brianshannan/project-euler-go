package main

import (
	"fmt"

	"utils"
)

func digit_fact_sum(n int) int {
	sum := 0
	for n != 0 {
		r := n % 10
		sum += utils.Factorial(r)
		n /= 10
	}
	return sum
}

func main() {
	sum := 0

	// max is all 9's. As 8*9! < 10^8, 10^8 provides a somewhat reasonable uper bound
	for i := 3; i < 10000000; i++ {
		if digit_fact_sum(i) == i {
			sum += i
		}
	}

	fmt.Println(sum)
}
