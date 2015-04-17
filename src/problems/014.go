package main

import "fmt"

func collatz_len(n int) int {
	count := 0

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count++
	}

	return count
}

func main() {
	max_start := 1000000
	max_len := 0
	max_len_num := 0

	for i := 2; i < max_start; i++ {
		seq_len := collatz_len(i)
		if seq_len > max_len {
			max_len = seq_len
			max_len_num = i
		}
	}

	fmt.Println(max_len_num)
}
