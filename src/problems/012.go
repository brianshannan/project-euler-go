package main

import (
	"fmt"
	"math"
)

func get_div_count(num int) int {
	count := 0
	for i := 1; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			count += 2
		}
	}
	return count
}

func main() {
	curr_num := 0
	for i := 1; ; i++ {
		curr_num += i
		if get_div_count(curr_num) > 500 {
			fmt.Println(curr_num)
			return
		}
	}
}
