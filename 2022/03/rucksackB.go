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

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var itemGroup []string
	var prioSum int
	for scanner.Scan() {

		x := scanner.Text()
		item := x
		itemGroup = append(itemGroup, item)

		if len(itemGroup) == 3 {
			sharedType := compareItem(itemGroup)
			prio := calcPrio(sharedType)
			prioSum = prioSum + prio
			itemGroup = nil
		}
	}

	fmt.Println(prioSum)

	file.Close()
}

func compareItem(itemGroup []string) string {
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
