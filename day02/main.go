package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var inputSlice []string

func init() {
	input = strings.TrimSuffix(input, "\n")
	inputSlice = strings.Split(input, ",")
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	total := 0

	for _, v := range inputSlice {
		firstLastSlice := strings.Split(v, "-")
		firstValue, _ := strconv.Atoi(firstLastSlice[0])
		lastValue, _ := strconv.Atoi(firstLastSlice[1])
		for i := firstValue; i <= lastValue; i++ {
			if t := hasSequence(i, 2); t {
				total = total + i
			}
		}
	}

	fmt.Printf("Part One Total: %d\n", total)
}

func partTwo() {
	total := 0

	for _, v := range inputSlice {
		firstLastSlice := strings.Split(v, "-")
		firstValue, _ := strconv.Atoi(firstLastSlice[0])
		lastValue, _ := strconv.Atoi(firstLastSlice[1])
		for i := firstValue; i <= lastValue; i++ {
			switch {
			case hasSequence(i, 2):
				total = total + i
			case hasSequence(i, 3):
				total = total + i
			case hasSequence(i, 5):
				total = total + i
			case hasSequence(i, 7):
				total = total + i
			}
		}
	}

	fmt.Printf("Part Two Total: %d\n", total)
}

func hasSequence(s int, numSequences int) bool {
	strVal := strconv.Itoa(s)
	strLen := len(strVal)

	// if the value's number of digits is not evenly divisible
	// by the number of sequences, it cannot have many sequences
	if strLen%numSequences != 0 {
		return false
	}

	seqSize := strLen / numSequences

	// for a sequence k to repeat multiple times in an integer, you
	// left shift it and fill it in with itself. For each repetition,
	// it is left shifted by its own length. This can be reversed by
	// dividing the resulting value by 10 left shifted the same number
	// of times, plus 1 for filling itself into the zeroes. This is
	// expressed mathematically for sequence k occurring n times as:
	// divisor = 10^(k*(n-1))+10^(k*(n-2))+...+10^k+1
	var divisor int
	for i := numSequences - 1; i > 0; i-- {
		divisor = divisor + int(math.Pow10(seqSize*i))
	}

	divisor = divisor + 1

	if s%divisor == 0 {
		return true
	}
	return false
}
