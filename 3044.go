func mostFrequentPrime(mat [][]int) int {
    primeNumbersCount := map[int]int{}
    for i := 0; i < len(mat); i++{
        for j := 0 ; j < len(mat[i]); j++{
            newPrimeNumbers := findPrimeNumbers(mat, i, j)
            for _, num := range newPrimeNumbers{
                primeNumbersCount[num]++
            }
        }
    }
    maxFreaquency := 0
    primeNumber := -1
    for k := range primeNumbersCount{
        if maxFreaquency < primeNumbersCount[k] || maxFreaquency == primeNumbersCount[k] && primeNumber < k{
            maxFreaquency = primeNumbersCount[k]
            primeNumber = k
        } 
    }
    return primeNumber
}
func findPrimeNumbers(mat [][]int, ii, jj int) []int{
    primeNumbers := []int{}
    // up
    number := 0
    i := ii
    j := jj
    for i >= 0{
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        i--
    }
    // down
    number = 0
    i = ii
    j = jj
    for i < len(mat){
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        i++
    }
    // left
    i = ii
    j = jj
    number = 0
    for j >= 0{
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        j--
    }
    // right
    i = ii
    j = jj
    number = 0
    for j < len(mat[0]){
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        j++
    }
    // upRight
    i = ii
    j = jj
    number = 0
    for j < len(mat[0]) && i >= 0{
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        i--
        j++
    }
    // upLeft
    i = ii
    j = jj
    number = 0
    for j >= 0 && i >= 0{
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        i--
        j--
    }
    // downRight
    i = ii
    j = jj
    number = 0
    for j < len(mat[0]) && i < len(mat){
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        i++
        j++
    }
    // downLeft
    i = ii
    j = jj
    number = 0
    for j >= 0 && i < len(mat){
        number = number * 10 + mat[i][j]
        if isPrime(number){
            primeNumbers = append(primeNumbers, number)
        }
        i++
        j--
    }

    return primeNumbers
}
func isPrime(num int) bool{
    if num == 1 || num < 10{
        return false
    }
    for i := 2; i < num / 2; i++{
        if num % i == 0{
            return false
        }
    }
    return true
}