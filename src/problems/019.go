package main

import "fmt"

// const days_with_leap = 365
// const days_without_leap = 366
//
// func days_in_year(n int) int {
//     if n % 4 == 0  {
//         if n % 100 == 0 && n % 400 != 0 {
//             return days_without_leap
//         }
//         return days_with_leap
//     }
//     return days_without_leap
// }

func days_in_month(month int, year int) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 1 && year%4 == 0 && !(year%100 == 0 && year%400 != 0) {
		return months[1] + 1
	}
	return months[month]
}

func main() {
	year := 1900
	month := 0
	day := 0
	count := 0

	for year <= 2000 {
		if year >= 1901 && day == 6 {
			count++
		}
		day += days_in_month(month, year)
		day %= 7
		year += (month + 1) / 12
		month = (month + 1) % 12
	}

	fmt.Println(count)
}
