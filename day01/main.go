package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSuffix(input, "\n")
	input = strings.ReplaceAll(input, "R", "")
	input = strings.ReplaceAll(input, "L", "-")
	dial := 50
	count := 0
	inputSlice := strings.SplitSeq(input, "\n")

	for v := range inputSlice {
		// turn the dial per the value
		turn, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		// we only care if it lands on zero, with a wrapping range of 0..99
		// so mod 100 on the turn keeps us within bounds
		dial = (dial + turn) % 100

		// check if we're on 0 and increment the count
		if dial == 0 {
			count++
		}
	}

	fmt.Printf("part one solution: %d", count)
}
