package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveB(fileContent []string) {
	zeros := [][]int{}
	grid := make([][]int, len(fileContent))
	for i, line := range fileContent {
		numarray := make([]int, len(line))
		for j, numstr := range strings.Split(line, "") {
			num, _ := strconv.Atoi(numstr)
			numarray[j] = num
			if num == 0 {
				zeros = append(zeros, []int{j, i})
			}
		}
		grid[i] = numarray
	}

	nines := 0
	var uphill func(x int, y int)
	uphill = func(x int, y int) {
		if grid[y][x] == 9 {
			nines++
			return
		}
		if y-1 >= 0 {
			if grid[y-1][x] == grid[y][x]+1 {
				uphill(x, y-1)
			}
		}
		if y+1 < len(grid) {
			if grid[y+1][x] == grid[y][x]+1 {
				uphill(x, y+1)
			}
		}
		if x-1 >= 0 {
			if grid[y][x-1] == grid[y][x]+1 {
				uphill(x-1, y)
			}
		}
		if x+1 < len(grid[0]) {
			if grid[y][x+1] == grid[y][x]+1 {
				uphill(x+1, y)
			}
		}
		return
	}

	for _, zero := range zeros {
		uphill(zero[0], zero[1])
	}

	fmt.Println(nines)
}
