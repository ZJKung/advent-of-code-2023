package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	sum := partOne(lines)
	fmt.Println(sum)
	sum = partTwo(lines)
	fmt.Println(sum)
}

func partOne(lines []string) int {
	sum := 0
	for _, line := range lines {
		digits := []int{}
		for _, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, digit)
			}
		}
		if len(digits) > 0 {
			first, last := digits[0], digits[len(digits)-1]
			sum += first*10 + last
		}
	}
	return sum
}
func partTwo(lines []string) int {
	spellingMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	sum := 0
	for _, line := range lines {
		digits := []int{}
		for i, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, digit)
			} else {
				for spelling, number := range spellingMap {
					if strings.HasPrefix(line[i:], spelling) {
						digits = append(digits, number)
						break
					}
				}
			}
		}
		if len(digits) > 0 {
			first, last := digits[0], digits[len(digits)-1]
			sum += first*10 + last
		}
	}
	return sum
}
