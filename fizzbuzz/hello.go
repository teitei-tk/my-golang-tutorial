package main

import (
	"fmt"
	"strconv"
)

const (
	NewLine = "\n"
)

func main() {
	threeDivision := "hoge"
	fiveDivision := "foo"

	maxLength := 50
	for i := 1; i <= maxLength; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf(threeDivision + fiveDivision)

		} else if i%3 == 0 {
			fmt.Printf(threeDivision)

		} else if i%5 == 0 {
			fmt.Printf(fiveDivision)

		} else {
			fmt.Printf(strconv.Itoa(i))
		}

		fmt.Printf(NewLine)
	}
}
