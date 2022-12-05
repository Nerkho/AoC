package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	fmt.Println(err)
	return
}

var stacks map[int][]string

func main() {

	stacks = map[int][]string{
		1: {"M", "S", "J", "L", "V", "F", "N", "R"},
		2: {"H", "W", "J", "F", "Z", "D", "N", "P"},
		3: {"G", "D", "C", "R", "W"},
		4: {"S", "B", "N"},
		5: {"N", "F", "B", "C", "P", "W", "Z", "M"},
		6: {"W", "M", "R", "P"},
		7: {"W", "S", "L", "G", "N", "T", "R"},
		8: {"V", "B", "N", "F", "H", "T", "Q"},
		9: {"F", "N", "Z", "H", "M", "L"},
	}

	file, err := os.Open("input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		move(moves, stacks)
	}
	for i := 1; i < len(stacks)+1; i++ {
		fmt.Println(stacks[i][0])
	}

}

func move(moves []string, stacks map[int][]string) {
	amnt, err := strconv.Atoi(moves[1])
	posA, err := strconv.Atoi(moves[3])
	posB, err := strconv.Atoi(moves[5])
	check(err)
	for i := amnt - 1; i >= 0; i-- {
		toMove := stacks[posA][i]
		remove(stacks[posA], i)
		stacks[posB] = append([]string{toMove}, stacks[posB]...)
	}

}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
