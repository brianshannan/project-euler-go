package main

import "fmt"

func decimal_cycle_len(n int) int {
	// No cycle if 2 or 5 divide n as they are factors of 10
	if n%2 == 0 || n%5 == 0 {
		return 0
	}

	// find minimum a s.t. 10^a == 1 mod n where a > 0
	num := 1
	count := 0
	for {
		num *= 10
		num %= n
		count++

		if num == 1 {
			return count
		}
	}
}

func main() {
	max_num := 1000
	max_len := 0
	num := 0

	for i := 2; i < max_num; i++ {
		cl := decimal_cycle_len(i)
		if cl > max_len {
			max_len = cl
			num = i
		}
	}

	fmt.Println(num)
}
