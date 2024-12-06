package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.in")
	fileContent := string(content)
	sections := strings.Split(fileContent, "\r\n\r\n")
	rules := sections[0]
	pages := sections[1]

	// initialize a map with key of [int] and value of []int slice
	dict := make(map[int][]int)
	for _, nums := range strings.Fields(rules) {
		num1, _ := strconv.Atoi(strings.Split(nums, "|")[0])
		num2, _ := strconv.Atoi(strings.Split(nums, "|")[1])
		dict[num1] = append(dict[num1], num2)
	}

	total := 0
	for _, line := range strings.Fields(pages) {
		isValid := true
		updates := strings.Split(line, ",")
		for i := len(updates) - 1; i >= 0; i-- {
			ruleNum, _ := strconv.Atoi(updates[i])
			for _, ele := range updates[:i] {
				num, _ := strconv.Atoi(ele)
				for _, val := range dict[ruleNum] {
					if num == val {
						isValid = false
						break
					}
				}
			}
		}
		if isValid {
			mid, _ := strconv.Atoi(updates[len(updates)/2])
			total += mid
		}
	}

	fmt.Println(total)
}
