package day1

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

	path, err := filepath.Abs("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rawData, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	content := string(rawData)

	values := strings.Split(content, "\n")

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

		// if len(intDigits) == 0 {
		// 	continue
		// }

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
