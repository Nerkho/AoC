package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var prioSumA int
var prioSumB int
var itemGroup []string

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
	for scanner.Scan() {
		x := scanner.Text()
		prioSumA = prioSumA + partA(x)
		prioSumB = prioSumB + partB(x)
	}
	fmt.Println("Part A :", prioSumA)
	fmt.Println("Part B :", prioSumB)
}

func partA(x string) int {
	var prio int
	item1 := x[0 : len(x)/2]
	item2 := x[len(x)/2:]
	sharedType := compareItemA(item1, item2)
	prio = calcPrio(sharedType)
	return prio
}

func compareItemA(item1 string, item2 string) string {
	var sharedType string
	c := []rune(item1)
	for i := 0; i < len(item1); i++ {
		c := string(c[i])
		if strings.Contains(item2, c) {
			sharedType = c
		}
	}
	return sharedType
}

func partB(item string) int {
	itemGroup = append(itemGroup, item)
	var prio int
	if len(itemGroup) == 3 {
		sharedType := compareItemB(itemGroup)
		prio = calcPrio(sharedType)
		itemGroup = nil
	}
	return prio
}

func compareItemB(itemGroup []string) string {
	var sharedType string
	x := []int{len(itemGroup[0]), len(itemGroup[1]), len(itemGroup[2])}
	bigItemSize := findMax(x)
	var bigItem string
	for i := 0; i < len(itemGroup); i++ {
		if bigItemSize == x[i] {
			bigItem = itemGroup[i]
		}
	}
	c := []rune(bigItem)
	for i := 0; i < bigItemSize; i++ {
		c := string(c[i])
		if strings.Contains(itemGroup[0], c) && strings.Contains(itemGroup[1], c) && strings.Contains(itemGroup[2], c) {
			sharedType = c
		}
	}
	return sharedType
}

func findMax(x []int) int {
	max := x[0]
	for _, y := range x {
		if y > max {
			max = y
		}
	}
	return max
}

func calcPrio(itemType string) int {
	var prio int

	prioList := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(prioList); i++ {
		if itemType == string(prioList[i]) {
			prio = i + 1
			break
		}
	}
	return prio
}
