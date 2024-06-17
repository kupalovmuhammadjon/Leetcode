func judgeSquareSum(c int) bool {
    limit := int(math.Sqrt(float64(c)))
    for i := 0; i <= limit; i++{
        if isSquare(c-i*i){
            return true
        }
    }
    return false
}

func isSquare(x int) bool{
    sqrt := int(math.Sqrt(float64(x)))
    if sqrt * sqrt == x{
        return true
    }
    return false
}