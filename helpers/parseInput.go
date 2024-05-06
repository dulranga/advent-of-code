package helpers

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ParseInputFile(relativePathFromRoot string) []string {

	path, err := filepath.Abs(relativePathFromRoot)
	if err != nil {
		log.Fatal(err)
	}

	rawData, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	content := string(rawData)

	values := strings.Split(content, "\n")

	return values
}
