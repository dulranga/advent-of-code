package main

import (
	"testing"

	challenges2024 "github.com/dulranga/advent-of-code/2024"
)

var left, right = challenges2024.GetSortedLists("../2024/inputs/day1.txt")

func BenchmarkDay1Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		challenges2024.CalculateDistancePart2(left, right)

	}

}

func BenchmarkDay1Part2Basic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		challenges2024.CalculateDistancePart2Basic(left, right)
	}
}
