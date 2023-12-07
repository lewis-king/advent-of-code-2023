package common

import (
	"log"
	"os"
	"strings"
)

func ReadFileLines(inputPath string) []string {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	// read lines from file
	return strings.Split(string(content), "\n")
}

func A() {

}
