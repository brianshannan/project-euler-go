package utils

func GetPrimesList(max_num int) []int {
    nums := make([]bool, max_num)
    for i := 0; i < len(nums); i++ {
        nums[i] = true
    }

    num_primes := 0
    for i := 2; i < max_num; i++ {
        if nums[i] == true {
            num_primes++
            for j := i * i; j < max_num; j += i {
                nums[j] = false
            }
        }
    }

    primes := make([]int, max_num)
    count := 0
    for i := 2; i < len(nums); i++ {
        if nums[i] == true {
            primes[count] = i
            count++
        }
    }
    return primes[:count]
}

func IsPrime(num int, primes []int) bool {
    for i := 0; i < len(primes) && primes[i] < num; i++ {
        if num%primes[i] == 0 {
            return false
        }
    }
    return true
}
