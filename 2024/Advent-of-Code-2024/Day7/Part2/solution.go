package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	content, _ := os.ReadFile("input.in")
	fileContent := strings.Split(string(content), "\r\n")
	// day7_2(string(content))
	solveA(fileContent)
	solveB(fileContent)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func solveB(fileContent []string) {
	total := 0
	for _, line := range fileContent {
		test, _ := strconv.Atoi(strings.Split(line, ":")[0])
		numstr := strings.Fields(strings.Split(line, ":")[1])
		nums := make([]int, len(numstr))

		for i, num := range numstr {
			nums[i], _ = strconv.Atoi(num)
		}

		if checkEquation(nums, test) {
			total += test
		}
	}
	fmt.Println(total) //348360680516005
}

func checkEquation(nums []int, test int) bool {
	ops := []string{"+", "*", "|"}
	combinations := make([][]string, 0)

	// Initial combinations with first operation
	for _, op := range ops {
		combinations = append(combinations, []string{op})
	}

	// Generate all possible combinations of operations
	for i := 1; i < len(nums)-1; i++ {
		newCombinations := make([][]string, 0)
		for _, combination := range combinations {
			for _, op := range ops {
				newCombination := make([]string, len(combination)+1)
				copy(newCombination, combination)
				newCombination[len(combination)] = op
				newCombinations = append(newCombinations, newCombination)
			}
		}
		combinations = newCombinations
	}

	for _, combination := range combinations {
		res := nums[0]
		for i := 0; i < len(combination); i++ {
			if combination[i] == "+" {
				res += nums[i+1]
			} else if combination[i] == "*" {
				res *= nums[i+1]
			} else if combination[i] == "|" {
				concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", res, nums[i+1]))
				res = concat
			}
		}
		if res == test {
			return true
		}
	}
	return false
}

func solveA(fileContent []string) {

	var eval func(test int, nums []int, idx int, res int) bool
	eval = func(test int, nums []int, idx int, res int) bool {

		if idx == len(nums) {
			return res == test
		}
		concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", res, nums[idx]))

		return eval(test, nums, idx+1, concat) || eval(test, nums, idx+1, res+nums[idx]) || eval(test, nums, idx+1, res*nums[idx])
	}

	total := 0

	for _, line := range fileContent {
		test, _ := strconv.Atoi(strings.Split(line, ":")[0])
		numstr := strings.Fields(strings.Split(line, ":")[1])
		nums := make([]int, len(numstr))

		for i, num := range numstr {
			nums[i], _ = strconv.Atoi(num)
		}

		if eval(test, nums, 0, 0) {
			total += test
		}
	}

	fmt.Println(total)
}
