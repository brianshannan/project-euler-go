package main

import (
	"fmt"
	"math"
)

func amicable(n int) int {
	sum := 1
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	return sum
}

func main() {
	max := 10000
	nums := make([]int, max)

	for i := 2; i < max; i++ {
		nums[i] = amicable(i)
	}

	sum := 0
	for i := 2; i < max; i++ {
		if nums[i] < max && i == nums[nums[i]] && i != nums[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}
