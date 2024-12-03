package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(a))
		for i, v := range a {
			nums[i], err = strconv.Atoi(v)
			if err != nil {
				fmt.Println("Conversion error !")
				return
			}
		}
		if nums[len(nums)-1] < nums[0] {
			nums = reverse(nums)
		}
		total += safeCheck(nums)
	}
	fmt.Println(total)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func safeCheck(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 3 || nums[i]-nums[i-1] < 1 {
			return 0
		}
	}
	return 1
}
