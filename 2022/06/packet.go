package main

import (
	"bufio"
	"fmt"
	"os"
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
	var x string
	for scanner.Scan() {
		x = scanner.Text()
	}
	a := common(x, 4)
	b := common(x, 14)
	fmt.Println("Part A :", a)
	fmt.Println("Part B :", b)
}

func common(input string, length int) int {
	numberChar := 0
	for i := length - 1; i < len(input); i++ {
		m := map[byte]int{}
		for j := 0; j < length; j++ {
			m[input[i-j]]++
		}
		if len(m) == length {
			numberChar = i + 1
			break
		}
	}
	return numberChar
}
