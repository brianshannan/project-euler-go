package main

import (
	"fmt"

	"utils"
)

func is_palindrome_base10(n int) bool {
	s := fmt.Sprintf("%d", n)
	r := utils.Reverse(s)
	return s == r
}

func is_palindrome_base2(n int) bool {
	s := fmt.Sprintf("%b", n)
	r := utils.Reverse(s)
	return s == r
}

func main() {
	max := 1000000
	sum := 0

	// Can't be even as there is a trailing zero in base 2
	// Could refine to make msd in base 10 odd as well,
	// but this runs fast enough
	for i := 1; i < max; i += 2 {
		if is_palindrome_base10(i) && is_palindrome_base2(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
