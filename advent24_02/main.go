package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	safeLineCount, bufferedSafeLineCount, err := countSafeReadings(reader)
	if err != nil {
		panic(fmt.Sprintf("Failed to count safe lines: %s", err.Error()))
	}

	fmt.Printf("Number of safe readings: %d\n", safeLineCount)
	fmt.Printf("Number of safe readings with buffer: %d\n", bufferedSafeLineCount)
}

func countSafeReadings(reader *bufio.Reader) (int, int, error) {
	var safeLineCount int = 0
	var bufferedSafeLineCount int = 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return safeLineCount, bufferedSafeLineCount, nil
		}
		if err != nil {
			return safeLineCount, bufferedSafeLineCount, err
		}

		values := convertLine(string(line))
		if areValuesSafe(values) {
			safeLineCount += 1
		}
		if areValuesSafeBuffered(values) {
			bufferedSafeLineCount += 1
		}
	}
}

func areValuesSafe(values []int) bool {
	if len(values) < 2 {
		return true
	}
	ascending := values[0] < values[1]
	for i := range len(values) - 1 {
		prev := values[i]
		cur := values[i+1]

		if prev < cur != ascending {
			return false
		}
		dist := distance(prev, cur)
		if dist < 1 || dist > 3 {
			return false
		}
	}
	return true
}

func areValuesSafeBuffered(values []int) bool {
	if areValuesSafe(values) {
		return true
	}

	for i := range values {
		valueCopy := make([]int, 0, len(values)-1)
		valueCopy = append(valueCopy, values[:i]...)
		valueCopy = append(valueCopy, values[i+1:]...)
		if areValuesSafe(valueCopy) {
			return true
		}
	}

	return false
}

func convertLine(line string) []int {
	parts := strings.Split(line, " ")
	ints := make([]int, len(parts))
	for i, part := range parts {
		parsed, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse line - invalid syntax: '%s'", line))
		}
		ints[i] = int(parsed)
	}
	return ints
}

func distance(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
