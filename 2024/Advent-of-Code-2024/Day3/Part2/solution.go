package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.in")
	// file, err := os.Open("test.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var str string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text()
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(str, -1)

	total := 0
	isEnable := true
	for _, match := range matches {
		if match[0] == "do()" {
			isEnable = true
		} else if match[0] == "don't()" {
			isEnable = false
		} else {
			if isEnable {
				total += multiply(match)
			}
		}
	}
	fmt.Println(total)
}

func multiply(str []string) int {
	num1, _ := strconv.Atoi(str[1])
	num2, _ := strconv.Atoi(str[2])
	return num1 * num2
}
