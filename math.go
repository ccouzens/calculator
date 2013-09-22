package main

import (
	"fmt"
	"strconv"
)

func parse(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}

func main() {
	input := "43"
	output, err := parse(input)
	if err == nil {
		fmt.Printf("%s = %d\n", input, output)
	} else {
		fmt.Printf("%s resulted in %s\n", input, err.Error())
	}
}
