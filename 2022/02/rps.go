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
		result := score1(x)
		scoreA = scoreA + result
		result = score2(x)
		scoreB = scoreB + result
	}

	fmt.Println("Part A :", scoreA)
	fmt.Println("Part B :", scoreB)
}

func score1(game []string) int {
	var score int

	switch game[0] {
	case "A":
		switch game[1] {
		case "X":
			score = 4
		case "Y":
			score = 8
		case "Z":
			score = 3
		}

	case "B":
		switch game[1] {
		case "X":
			score = 1
		case "Y":
			score = 5
		case "Z":
			score = 9
		}

	case "C":
		switch game[1] {
		case "X":
			score = 7
		case "Y":
			score = 2
		case "Z":
			score = 6
		}
	}

	return score
}

func score2(game []string) int {
	var score int

	switch game[0] {
	case "A":
		switch game[1] {
		case "X":
			score = 3
		case "Y":
			score = 4
		case "Z":
			score = 8
		}

	case "B":
		switch game[1] {
		case "X":
			score = 1
		case "Y":
			score = 5
		case "Z":
			score = 9
		}

	case "C":
		switch game[1] {
		case "X":
			score = 2
		case "Y":
			score = 6
		case "Z":
			score = 7
		}
	}

	return score
}
