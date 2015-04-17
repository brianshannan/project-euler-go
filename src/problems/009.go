package main

import "fmt"

func main() {
	target := 1000

	for c := 1; c < target; c++ {
		for b := 1; b < c; b++ {
			for a := 1; a < b; a++ {
				if a+b+c == target && a*a+b*b == c*c {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}
