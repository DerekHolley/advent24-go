package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}
	contents, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}

	reg, err := regexp.Compile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	if err != nil {
		panic(fmt.Sprintf("Failed to compile regex: %s", err.Error()))
	}

	subreg, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(fmt.Sprintf("Failed to compile subregex: %s", err.Error()))
	}

	matches := reg.FindAllString(string(contents), -1)

	var sum int64 = 0
	var doing bool = true
	for _, match := range matches {
		if match == "do()" {
			doing = true
			continue
		}
		if match == "don't()" {
			doing = false
			continue
		}
		if !doing {
			continue
		}
		parts := subreg.FindAllString(match, -1)
		if len(parts) != 2 {
			panic(fmt.Sprintf("Failed to parse line '%s'", match))
		}
		lhs, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse integer part '%s': %s", parts[0], err.Error()))
		}

		rhs, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse integer part '%s': %s", parts[1], err.Error()))
		}
		sum += lhs * rhs
	}
	fmt.Printf("Total sum of valid expressions: %d", sum)
}
