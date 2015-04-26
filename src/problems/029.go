package main

import (
	"fmt"
	"math/big"
)

func find(n *big.Int, nums [][]*big.Int, curr int) bool {
	for i := 0; i < curr; i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == nil {
				return false
			} else if nums[i][j].Cmp(n) == 0 {
				return true
			} else if nums[i][j].Cmp(n) == 1 {
				break
			}
		}
	}
	return false
}

func main() {
	max := 100
	nums := make([][]*big.Int, max-1)
	for i, _ := range nums {
		nums[i] = make([]*big.Int, max-1)
	}
	c := 0

	for a := 2; a <= max; a++ {
		for b := 2; b <= max; b++ {
			n := big.NewInt(int64(a))
			n.Exp(n, big.NewInt(int64(b)), nil)
			nums[a-2][b-2] = n
			if !find(n, nums, a-2) {
				c++
			}
		}
	}

	fmt.Println(c)
}
