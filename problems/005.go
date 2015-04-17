package main

import "fmt"

func gcd(a int, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(b-a, a)
	}
}

func main() {
	curr_lcm := 1
	for i := 2; i <= 20; i++ {
		curr_lcm = (curr_lcm * i) / gcd(curr_lcm, i)
	}
	fmt.Println(curr_lcm)
}
