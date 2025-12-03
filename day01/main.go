package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var inputStringSlice []string
var inputIntSlice []int
var startingDial = 50

func init() {
	input = strings.TrimSuffix(input, "\n")
	input = strings.ReplaceAll(input, "R", "")
	input = strings.ReplaceAll(input, "L", "-")
	inputStringSlice = strings.Split(input, "\n")
	inputIntSlice = make([]int, len(inputStringSlice))
	for i, s := range inputStringSlice {
		inputIntSlice[i], _ = strconv.Atoi(s)
	}
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	count := 0
	dial := startingDial

	for _, turn := range inputIntSlice {
		// we only care if it lands on zero, with a wrapping range of 0..99
		// so mod 100 on the turn keeps us within bounds
		dial = (dial + turn) % 100

		// check if we're on 0 and increment the count
		if dial == 0 {
			count++
		}
	}

	fmt.Printf("part one solution: %d\n", count)
}

func partTwo() {
	count := 0
	dial := startingDial

	for _, turn := range inputIntSlice {
		for range Abs(turn) {
			if turn > 0 {
				dial = (dial + 1) % 100
			} else {
				dial = (dial - 1) % 100
			}
			if dial == 0 {
				count++
			}
		}

	}

	fmt.Printf("part two solution: %d\n", count)
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
