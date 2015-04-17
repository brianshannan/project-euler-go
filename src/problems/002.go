package main

import "fmt"

func main() {
	num1 := 1
	num2 := 1
	sum := 0

	for num2 < 4000000 {
		if num2%2 == 0 {
			sum += num2
		}

		temp := num1 + num2
		num1 = num2
		num2 = temp
	}

	fmt.Println(sum)
}
