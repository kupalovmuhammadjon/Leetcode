func maxBottlesDrunk(numBottles int, numExchange int) int {
    emptyBottles := numBottles
    numberOfBottlesDrank := numBottles
    numBottles = 0
    
    for emptyBottles >= numExchange{
        for emptyBottles >= numExchange{
            numBottles++
            emptyBottles -= numExchange
            numExchange++
        }
        numberOfBottlesDrank += numBottles
        emptyBottles += numBottles
        numBottles = 0
    }
    return numberOfBottlesDrank
}