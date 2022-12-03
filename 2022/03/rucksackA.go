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

	var prioSum int

	for scanner.Scan() {
		x := scanner.Text()
		item1 := x[0 : len(x)/2]
		item2 := x[len(x)/2:]
		sharedType := compareItem(item1, item2)
		prio := calcPrio(sharedType)
		prioSum = prioSum + prio
	}
	fmt.Println(prioSum)
}

func compareItem(item1 string, item2 string) string {
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
