package main

import (
	"fmt"
	"math"
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
	red := 12
	green := 13
	blue := 14

	for i, str := range lines {
		ID := i + 1
		game_data := strings.Split(str, ":")[1]
		games := strings.Split(game_data, ";")
		failed := false
		for _, game := range games {
			if checkColorExceedLimit(game, "blue", blue) {
				failed = true
			}
			if checkColorExceedLimit(game, "red", red) {
				failed = true
			}
			if checkColorExceedLimit(game, "green", green) {
				failed = true
			}
		}
		if !failed {
			sum += ID
		}
	}
	return sum
}
func partTwo(lines []string) int {
	sum := 0
	for _, str := range lines {
		game_data := strings.Split(str, ":")[1]
		games := strings.Split(game_data, ";")
		m := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		mul := 1
		for _, game := range games {
			for k, v := range m {
				m[k] = int(math.Max(float64(v), float64(getNum(game, k))))
			}
		}
		for _, v := range m {
			mul *= v
		}
		sum += mul
	}
	return sum
}
func checkColorExceedLimit(game string, color string, limit int) bool {
	return getNum(game, color) > limit
}
func getNum(game string, color string) int {
	if idx := strings.Index(game, color); idx != -1 {
		num, _ := strconv.Atoi(strings.TrimSpace(string(game[(idx - 3):idx])))
		return num
	}
	return 0
}
