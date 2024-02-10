package main

import (
	"fmt"
	"math"
)

func countTestedDevices(batteryPercentages []int) int {
	count := 0
	tested := make(map[int]bool)

	for i := 0; i < len(batteryPercentages); i++ {
		absoluteValue := int(math.Abs(float64(batteryPercentages[i])))
		if !tested[absoluteValue] {
			count++
			tested[absoluteValue] = true
		}
	}

	return count
}

func main() {
	batteryPercentages := []int{2, 1}
	result := countTestedDevices(batteryPercentages)
	fmt.Println("Result:", result)
}
