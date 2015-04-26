package main

import "fmt"

func calc_rec(coins []int, min_index int, val int) int {
	if val < 0 {
		return 0
	} else if val == 0 {
		return 1
	}

	count := 0
	for i := min_index; i < len(coins); i++ {
		count += calc_rec(coins, i, val-coins[i])
	}
	return count
}

func calc_dp(coins []int, val int) int {
	counts := make([][]int, val+1)
	for i := 0; i < val+1; i++ {
		counts[i] = make([]int, len(coins))
	}
	for i := 0; i < len(coins); i++ {
		counts[0][i] = 1
	}

	for i := 1; i < val+1; i++ {
		for j := 0; j < len(coins); j++ {
			for k := 0; k <= j; k++ {
				if i-coins[k] >= 0 {
					counts[i][j] += counts[i-coins[k]][k]
				}
			}
		}
	}

	return counts[val][len(coins)-1]
}

func main() {
	max := 200
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	fmt.Println(calc_dp(coins, max))
}
