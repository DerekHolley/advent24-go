package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Extract single line from buffered reader
func getLine(r *bufio.Reader) (string, error) {
	bytes, _, err := r.ReadLine()
	return string(bytes), err
}

// Parse a line (splitting on "   " - three spaces) into two integers
func parseLine(str string) (int, int, error) {
	parts := strings.Split(str, "   ")
	if len(parts) != 2 {
		return -1, -1, errors.New("failed to parse line - split failed")
	}
	num1, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	num2, err := strconv.ParseInt(parts[1], 10, 64)
	return int(num1), int(num2), err
}

// Read the file to extract two lists of length 1000
func getLists() ([]int, []int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, nil, err
	}
	reader := bufio.NewReader(file)

	//we happen to know the file is 1000 lines long
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	//Read the file until it's empty, and return the values
	for {
		line, err := getLine(reader)
		if err == io.EOF {
			return left, right, nil
		}
		if err != nil {
			return nil, nil, err
		}
		leftVal, rightVal, err := parseLine(line)
		if err != nil {
			return nil, nil, err
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}
}

func absDifference(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func calculateSumOfDifference(left, right []int) int {
	var sum int = 0
	for i := range len(left) {
		sum += absDifference(left[i], right[i])
	}
	return sum
}

func calculateSimilarityScore(left, right []int) int {
	sum := 0
	for _, l := range left {
		for _, r := range right {
			if l == r {
				sum += l
			}
		}
	}
	return sum
}

func main() {
	left, right, err := getLists()
	if err != nil {
		panic("Failed to get lists: " + err.Error())
	}
	//Sort the slices, which will match up the left and right values
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sumOfDifferences := calculateSumOfDifference(left, right)
	similarityScore := calculateSimilarityScore(left, right)

	fmt.Printf("Sum of differences:\n%d\n", sumOfDifferences)
	fmt.Printf("Similarity Score:\n%d\n", similarityScore)
}
