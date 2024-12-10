package day1

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/dulranga/advent-of-code/helpers"
)

var spelledDigitRegex = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)

var DigitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Day1() {
	values := helpers.ParseInputFile("2023/day1/input.txt")
	sum := ParseWithSpelledDigits(values, spelledDigitRegex)
	fmt.Printf("Sum: %v\n", sum)
}

func ParseWithSpelledDigits(values []string, regex *regexp.Regexp) int64 {
	var sum int64
	for v := range values {
		alteredText := values[v]

		// TODO: Add logic to identify "twone" as two, one both

		// -1 to get all matches
		digits := regex.FindAllString(alteredText, -1)

		// pre allocation to avoid making new copies each time adding the mapped digit
		intDigits := make([]int, len(digits))

		for digit := range digits {
			// transformation of the digit to integer
			intDigits[digit] = DigitMap[digits[digit]]
		}

		if len(intDigits) == 0 {
			continue
		}

		firstDigit := intDigits[0]
		lastDigit := intDigits[len(intDigits)-1]

		// string concatenation to get the 2 digit number
		calibrationString := fmt.Sprintf("%d%d", firstDigit, lastDigit)

		calibrationValue, err := strconv.ParseInt(calibrationString, 10, 8)
		if err != nil {
			log.Fatal(err)
		}

		sum += calibrationValue

	}

	return sum

}
