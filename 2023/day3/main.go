package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/dulranga/advent-of-code/helpers"
)

func Day3() {
	var sum int
	specialSymbolCheck := regexp.MustCompile(`[^0-9.]`)
	digitCheck := regexp.MustCompile(`[0-9]+`)

	input := helpers.ParseInputFile("2023/day3/input.txt")

	for lineIdx, row := range input {
		found := digitCheck.FindAllStringIndex(row, -1)
		if found == nil {
			continue
		}

		for _, idx := range found {
			start := idx[0]
			end := idx[1]

			currentItem := row[start:end]

			closest := findClosest(input, lineIdx, start, end)

			if specialSymbolCheck.MatchString(closest) {
				currentItemInt, err := strconv.ParseInt(currentItem, 10, 16)
				if err != nil {
					log.Fatal(err)
				}
				sum += int(currentItemInt)

			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func findClosest(input []string, lineIdx int, start int, end int) string {

	var output []string

	currentLine := input[lineIdx]
	// currentItem := currentLine[start:end] // end is not included

	// before and after elements of the item
	beforeIdx := start - 1
	afterIdx := end + 1

	// top bottom rows
	topClose := lineIdx - 1
	bottonClose := lineIdx + 1

	// correction to avoid accessing slice overflows
	if beforeIdx < 0 {
		beforeIdx = 0
	}
	if afterIdx > len(currentLine) {
		afterIdx = len(currentLine)
	}

	if beforeIdx != start {
		currentBefore := string(currentLine[beforeIdx])
		output = append(output, currentBefore)
	}

	if afterIdx != end {
		// after index is not included in the current element
		// so the one before it is the after element
		currentAfter := string(currentLine[afterIdx-1])
		output = append(output, currentAfter)
	}

	if lineIdx > 0 {
		topLine := input[topClose][beforeIdx:afterIdx]
		output = append(output, topLine)
	}

	// lineIdx + 1 to avoid finding close elements below last row
	if (lineIdx + 1) < len(input) {
		bottomLine := input[bottonClose][beforeIdx:afterIdx]
		output = append(output, bottomLine)
	}

	return strings.Join(output, "")

}
