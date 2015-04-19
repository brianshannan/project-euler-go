package main

import (
	"fmt"
	"sort"
)

type permutation []byte

func (p permutation) Len() int {
	return len(p)
}

func (p permutation) Less(i int, j int) bool {
	return p[i] < p[j]
}

func (p permutation) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	n := 1000000
	p := permutation{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	max_val := byte(11)

	for i := 1; i < n; i++ {
		first := -1
		// Find rightmost character smaller than the characater after it
		for j := len(p) - 2; j >= 0; j-- {
			if p[j] < p[j+1] {
				first = j
				break
			}
		}

		// If one wasn't found there wasn't another permutation
		if first == -1 {
			fmt.Println("Out of permutations already")
			return
		}

		// Find smallest character after the "first" character which is larger than the "first"
		next := first
		small_val := max_val
		for j := first + 1; j < len(p); j++ {
			if p[j] > p[first] && p[j] < small_val {
				next = j
				small_val = p[j]
			}
		}

		// Swap them and sort after the original "first" character
		p[first], p[next] = p[next], p[first]
		sort.Sort(p[first+1:])
	}

	fmt.Println(p)
}
