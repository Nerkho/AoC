package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(err error) {
	fmt.Println(err)
	return
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var tCal []int
	thisElf := 0

	for _, x := range input {
		if x == "" {
			tCal = append(tCal, thisElf)
			thisElf = 0
		} else {
			cal, _ := strconv.Atoi(x)
			thisElf = thisElf + cal
		}
	}

	sort.Ints(tCal)
	fmt.Println("Answer A")
	fmt.Println(tCal[len(tCal)-1])

	t3 := 0
	for _, x := range tCal[len(tCal)-3:] {
		t3 = t3 + x
	}
	fmt.Println("Answer B")
	fmt.Println(t3)
}
