package day1

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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

	for v := range values {
		fmt.Printf("values[v]: %v\n", values[v])
	}
}
