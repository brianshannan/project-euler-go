package main

import "fmt"

// Only works if n < 1000
func int_to_string(n int) string {
	ones := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	result := ""

	msd := n / 100
	lsd := n % 100

	if msd > 0 {
		result = ones[msd] + " hundred"
		if lsd > 0 {
			result += " and "
		}
	}

	if 0 < lsd && lsd < 20 {
		result += ones[lsd]
	} else {
		result += tens[lsd/10]
		if lsd%10 > 0 {
			result += "-" + ones[lsd%10]
		}
	}

	return result
}

// String length disregarding spaces and hyphens
func get_string_len(s string) int {
	l := 0
	for _, v := range s {
		if v != ' ' && v != '-' {
			l++
		}
	}
	return l
}

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		sum += get_string_len(int_to_string(i))
	}

	// Get 1000
	sum += 11

	fmt.Println(sum)
}
