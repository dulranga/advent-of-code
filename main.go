package main

import (
	"fmt"

	challenges2024 "github.com/dulranga/advent-of-code/2024"
)

func main() {
	var left, right = challenges2024.GetSortedLists("2024/inputs/day1.txt")

	fmt.Printf("challenges2024.CalculateDistancePart2(left, right): %v\n", challenges2024.CalculateDistancePart2(left, right))

	fmt.Printf("challenges2024.CalculateDistancePart2(left, right): %v\n", challenges2024.CalculateDistancePart2(left, right))

	// day1.Day1()
	// day2.Day2()
	challenges2024.Day1("./2024/inputs/day1.txt")
}
