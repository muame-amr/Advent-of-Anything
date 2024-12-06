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
	indegree := make(map[int][]int)
	for _, nums := range strings.Fields(rules) {
		num1, _ := strconv.Atoi(strings.Split(nums, "|")[0])
		num2, _ := strconv.Atoi(strings.Split(nums, "|")[1])
		indegree[num1] = append(indegree[num1], num2)
	}

	total := 0
	for _, line := range strings.Fields(pages) {
		isValid := true
		updates := strings.Split(line, ",")
		for i := len(updates) - 1; i >= 0; i-- {
			ruleNum, _ := strconv.Atoi(updates[i])
			for _, ele := range updates[:i] {
				num, _ := strconv.Atoi(ele)
				for _, val := range indegree[ruleNum] {
					if num == val {
						isValid = false
						break
					}
				}
			}
		}
		if !isValid {

			ordered := []int{}
			validMap := make(map[int]bool)
			outdegree := make(map[int][]int)

			for _, i := range updates {
				num, _ := strconv.Atoi(i)
				validMap[num] = true

				for _, out := range indegree[num] {
					outdegree[out] = append(outdegree[out], num)
				}
			}

			/*
				Check for number without outdegree and put it in
				front. Then, delete it from original updates (validMap).
				Also update the value of outdegree by removing the number
				from its list.
			*/
			for len(validMap) > 0 {
				for num := range validMap {
					if len(outdegree[num]) == 0 {
						ordered = append(ordered, num)
						delete(validMap, num)

						for k, v := range outdegree {
							l := []int{}

							for _, el := range v {
								if el != num {
									l = append(l, el) // updates
								}
							}

							outdegree[k] = l
						}
					}
				}
			}
			mid := ordered[len(ordered)/2]
			total += mid
		}
	}
	fmt.Println(total)
}

// Kudos to Josh Ackland for this solution: https://www.youtube.com/watch?v=Jb0_sXyWb-o
