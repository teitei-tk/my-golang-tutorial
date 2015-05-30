package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NewLine = "\n"
)

func main() {
	var fp *os.File
	var err error

	filePath := "./testfile.txt"
	fp, err = os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) <= 0 {
			continue
		}

		fmt.Printf(text)
		fmt.Printf(NewLine)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
