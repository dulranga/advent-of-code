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

func Day1() {

	digitRegex := regexp.MustCompile(`\d`)

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

	// holds the sum of all calibrated values
	var sum int64

	for v := range values {
		alteredText := values[v]

		// -1 to get all matches
		digits := digitRegex.FindAllString(alteredText, -1)

		if len(digits) == 0 {
			continue
		}

		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]

		// string concatenation to get the 2 digit number
		calibrationString := firstDigit + lastDigit

		calibrationValue, err := strconv.ParseInt(calibrationString, 10, 8)
		if err != nil {
			log.Fatal(err)
		}

		sum += calibrationValue

	}

	fmt.Printf("Sum: %v\n", sum)
}
