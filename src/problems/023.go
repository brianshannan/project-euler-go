package main

import (
	"fmt"
	"math"
)

func is_abundant(n int) bool {
	sum := 1

	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			sum += i
			div := n / i
			if i != div {
				sum += div
			}
		}
	}

	return sum > n
}

func abundant(max int) []int {
	a := make([]int, max)
	count := 0

	for i := 1; i <= max; i++ {
		if is_abundant(i) {
			a[count] = i
			count++
		}
	}

	return a[:count]
}

func main() {
	max := 28123
	a := abundant(max)
	nums := make([]bool, max+1)

	for i := 0; i < len(nums); i++ {
		nums[i] = true
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			sum := a[i] + a[j]
			if sum > max {
				break
			}
			nums[sum] = false
		}
	}

	sum := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == true {
			sum += i
		}
	}
	fmt.Println(sum)
}
