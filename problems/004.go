package main

import (
	"fmt"
	"strconv"
)

func is_palindrome(num int) bool {
	s := strconv.Itoa(num)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	max := 0

	for i := 999; i > 0; i-- {
		for j := 999; j >= i; j-- {
			prod := i * j
			if is_palindrome(prod) {
				if prod > max {
					max = prod
				}
			}
		}
	}

	fmt.Println(max)
}
