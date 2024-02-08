func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
    delay := arrivalTime + delayedTime
    return delay % 24
}