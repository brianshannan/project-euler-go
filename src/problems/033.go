package main

import (
	"fmt"

	"utils"
)

func fracs_equiv(n1 int, d1 int, n2 int, d2 int) bool {
	if n1 > n2 {
		n1, n2 = n2, n1
		d1, d2 = d2, d1
	}
	if n2%n1 != 0 {
		return false
	}
	r := n2 / n1
	return d2%r == 0 && d1 == d2/r
}

func can_cancel_digits(n int, d int) bool {
	n1 := n % 10
	n2 := n / 10
	d1 := d % 10
	d2 := d / 10

	gcd := utils.Gcd(n, d)
	nr := n / gcd
	dr := d / gcd

	return (n1 == d1 && fracs_equiv(n2, d2, nr, dr)) ||
		(n1 == d2 && fracs_equiv(n2, d1, nr, dr)) ||
		(n2 == d1 && fracs_equiv(n1, d2, nr, dr)) ||
		(n2 == d2 && fracs_equiv(n1, d1, nr, dr))
}

func main() {
	num_prod := 1
	den_prod := 1

	// Just brute force, really small sample size
	// must be double digits
	for i := 11; i <= 98; i++ {
		// multiples of 10 are "trivial"
		if i%10 == 0 {
			continue
		}
		for j := i + 1; j <= 99; j++ {
			if can_cancel_digits(i, j) {
				num_prod *= i
				den_prod *= j
			}
		}
	}

	// reduce fraction
	fmt.Println(den_prod / utils.Gcd(num_prod, den_prod))
}
