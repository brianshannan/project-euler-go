package utils

func pandigital_help(n int, arr []int, count *int) bool {
    for n != 0 {
        val := n % 10
        if val == 0 || arr[val-1] > 0 {
            return false
        }
        arr[val-1]++
        *count++
        n /= 10
    }
    return true
}

func IsPandigital(nums ...int) bool {
    arr := make([]int, 9)
    count := 0

    for _, val := range nums {
        if !pandigital_help(val, arr, &count) {
            return false
        }
    }
    
    return count == 9
}
