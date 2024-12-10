package challenges2023

import (
	"fmt"

	"github.com/dulranga/advent-of-code/helpers"
)

func Day2() {

	checker := CheckRecord{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	records := helpers.ParseInputFile("2023/day2/input.txt")

	var sum int
	var totalPower int

	for _, recordStr := range records {

		record := ParseRecord(recordStr)
		isPossible := record.IsPossible(checker)

		if isPossible {
			sum += record.GameId
		}

		minCheckRecord := record.GetMinPossibleCheckRecord()
		power := minCheckRecord.Blue * minCheckRecord.Green * minCheckRecord.Red

		totalPower += power
	}

	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("totalPower: %v\n", totalPower)

}
