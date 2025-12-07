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
}

func partOne() {
	total := 0

	for _, v := range inputSlice {
		firstLastSlice := strings.Split(v, "-")
		firstValue, _ := strconv.Atoi(firstLastSlice[0])
		lastValue, _ := strconv.Atoi(firstLastSlice[1])
		for i := firstValue; i <= lastValue; i++ {
			strVal := strconv.Itoa(i)
			strLen := len(strVal)
			if strLen%2 != 0 {
				continue
			}

			mid := strLen / 2
			power := int(math.Pow10(mid)) + 1
			if i%power == 0 {
				total = total + i
			}
		}
	}

	fmt.Printf("Part One Total: %d\n", total)
}
