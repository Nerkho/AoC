package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	fmt.Println(err)
	return
}

var scoreA int
var scoreB int

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scoreA = 0
	scoreB = 0
	for scanner.Scan() {
		x := strings.Fields(scanner.Text())
		resultA, resultB := score(x)
		scoreA = scoreA + resultA
		scoreB = scoreB + resultB
	}
	fmt.Println("Part A :", scoreA)
	fmt.Println("Part B :", scoreB)
}

func score(game []string) (int, int) {
	var scoreA int
	var scoreB int

	switch game[0] {
	case "A":
		switch game[1] {
		case "X":
			scoreA = 4
			scoreB = 3
		case "Y":
			scoreA = 8
			scoreB = 4
		case "Z":
			scoreA = 3
			scoreB = 8
		}

	case "B":
		switch game[1] {
		case "X":
			scoreA = 1
			scoreB = 1
		case "Y":
			scoreA = 5
			scoreB = 5
		case "Z":
			scoreA = 9
			scoreB = 9
		}

	case "C":
		switch game[1] {
		case "X":
			scoreA = 7
			scoreB = 2
		case "Y":
			scoreA = 2
			scoreB = 6
		case "Z":
			scoreA = 6
			scoreB = 7
		}
	}

	return scoreA, scoreB
}
