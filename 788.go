func rotatedDigits(n int) int {
    count := 0
    for i := 0; i <= n; i++ {
        num := i
        valid := false

        for num > 0 {
            digit := num % 10
            
            if digit == 3 || digit == 4 || digit == 7 {
                valid = false
                break
            }
            if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
                valid = true
            }
            num /= 10
        }

        if valid {
            count++
        }
    }

    return count
}
