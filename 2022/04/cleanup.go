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

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var a int
	var b int
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		if comparePairA(pair) {
			a++
		}
		if comparePairB(pair) {
			b++
		}
	}
	fmt.Println("Part A :", a)
	fmt.Println("Part B :", b)
}

func comparePairA(pair []string) bool {
	var pairInPair bool
	left := fullRange(pair[0])
	right := fullRange(pair[1])
	if contains(left, right[0]) && contains(left, right[len(right)-1]) || contains(right, left[0]) && contains(right, left[len(left)-1]) {
		pairInPair = true
	}
	return pairInPair
}

func comparePairB(pair []string) bool {
	var pairInPair bool
	left := fullRange(pair[0])
	right := fullRange(pair[1])
	if overlap(left, right) || overlap(right, left) {
		pairInPair = true
	}
	return pairInPair
}

func fullRange(oneSide string) []string {
	var fullRange []string
	fRange := strings.Split(oneSide, "-")
	min, err := strconv.Atoi(fRange[0])
	max, err := strconv.Atoi(fRange[1])
	check(err)
	for i := min; i < max+1; i++ {
		fullRange = append(fullRange, strconv.Itoa(i))
	}
	return fullRange
}

func overlap(s, e []string) bool {
	var b bool
	for _, a := range e {
		if contains(s, a) {
			b = true
		}
	}
	return b
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
