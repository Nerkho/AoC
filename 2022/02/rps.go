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
	fmt.Println(scoreA)
	fmt.Println(scoreB)

	file.Close()
}

func score1(game []string) int {
	var score int

	if game[0] == "A" && game[1] == "X" {
		score = 4
	}
	if game[0] == "A" && game[1] == "Y" {
		score = 8
	}
	if game[0] == "A" && game[1] == "Z" {
		score = 3
	}
	if game[0] == "B" && game[1] == "X" {
		score = 1
	}
	if game[0] == "B" && game[1] == "Y" {
		score = 5
	}
	if game[0] == "B" && game[1] == "Z" {
		score = 9
	}
	if game[0] == "C" && game[1] == "X" {
		score = 7
	}
	if game[0] == "C" && game[1] == "Y" {
		score = 2
	}
	if game[0] == "C" && game[1] == "Z" {
		score = 6
	}

	return score
}

func score2(game []string) int {
	var score int

	if game[0] == "A" && game[1] == "X" {
		score = 3
	}
	if game[0] == "A" && game[1] == "Y" {
		score = 4
	}
	if game[0] == "A" && game[1] == "Z" {
		score = 8
	}
	if game[0] == "B" && game[1] == "X" {
		score = 1
	}
	if game[0] == "B" && game[1] == "Y" {
		score = 5
	}
	if game[0] == "B" && game[1] == "Z" {
		score = 9
	}
	if game[0] == "C" && game[1] == "X" {
		score = 2
	}
	if game[0] == "C" && game[1] == "Y" {
		score = 6
	}
	if game[0] == "C" && game[1] == "Z" {
		score = 7
	}

	return score
}
