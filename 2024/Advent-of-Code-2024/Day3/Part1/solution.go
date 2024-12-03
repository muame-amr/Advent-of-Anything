package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(str, -1)

	total := 0
	for _, match := range matches {
		match = strings.Replace(match, "(", " ", 1)
		match = strings.Replace(match, ",", " ", 1)
		match = strings.Replace(match, ")", " ", 1)
		num1, _ := strconv.Atoi(strings.Fields(match)[1])
		num2, _ := strconv.Atoi(strings.Fields(match)[2])
		total += num1 * num2
	}
	fmt.Println(total)
}
