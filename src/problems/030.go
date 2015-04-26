package main

import "fmt"

func fifth_power(n int) int {
	return n * n * n * n * n
}

func num_to_fifth_sum(n int) int {
	sum := 0
	for n != 0 {
		sum += fifth_power(n % 10)
		n /= 10
	}
	return sum
}

func main() {
	max := 1000000
	sum := 0

	for i := 2; i < max; i++ {
		if i == num_to_fifth_sum(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
