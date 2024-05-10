package day2

import (
	"fmt"
	"log"
	"strings"

	"github.com/dulranga/advent-of-code/helpers"
)

type CheckRecord struct {
	Red   int
	Blue  int
	Green int
}

type Record struct {
	Count  Count
	GameId int
}

/* Number of times balls picked in each round */
type Count struct {
	Red   []int
	Green []int
	Blue  []int
}

func (r *Record) AddCount(name string, count int) {
	switch name {
	case "blue":
		r.Count.Blue = append(r.Count.Blue, count)
	case "green":
		r.Count.Green = append(r.Count.Green, count)
	case "red":
		r.Count.Red = append(r.Count.Red, count)
	default:
		break
	}
}

func (r *Record) IsPossible(check CheckRecord) bool {

	for _, count := range r.Count.Blue {
		if count > check.Blue {
			return false
		}
	}
	for _, count := range r.Count.Green {
		if count > check.Green {
			return false
		}
	}
	for _, count := range r.Count.Red {
		if count > check.Red {
			return false
		}
	}

	return true
}

/* Get the minimum amount of balls in each color required to play this game */
func (r *Record) GetMinPossibleCheckRecord() CheckRecord {
	return CheckRecord{
		Blue:  helpers.GetMax(r.Count.Blue),
		Red:   helpers.GetMax(r.Count.Red),
		Green: helpers.GetMax(r.Count.Green),
	}
}

func ParseRecord(recordStr string) Record {

	splitted := strings.Split(recordStr, ":")

	var gameId int
	var record Record

	_, err := fmt.Sscanf(splitted[0], "Game %d", &gameId)
	if err != nil {
		log.Fatal(err)
	}

	record.GameId = gameId

	tries := strings.Split(splitted[1], ";")

	for _, tryStr := range tries {

		picks := strings.Split(tryStr, ",")
		for _, pickStr := range picks {

			var count int
			var name string
			_, err := fmt.Sscanf(pickStr, "%d %v", &count, &name)
			if err != nil {
				log.Fatal(err)
			}
			record.AddCount(name, count)

		}

	}

	return record
}
