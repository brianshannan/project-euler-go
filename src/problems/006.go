package main

import "fmt"

func main() {
	num := 100

	sum_sqrs := 0
	for i := 1; i <= num; i++ {
		sum_sqrs += i * i
	}

	sqr_sum := 0
	for i := 1; i <= num; i++ {
		sqr_sum += i
	}
	sqr_sum = sqr_sum * sqr_sum

	fmt.Println(sqr_sum - sum_sqrs)
}
