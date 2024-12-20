package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveA(fileContent []string) {
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

	nines := map[string]bool{}
	var uphill func(x int, y int)
	uphill = func(x int, y int) {
		if grid[y][x] == 9 {
			nines[fmt.Sprintf("%d,%d", x, y)] = true
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

	total := 0
	for _, zero := range zeros {
		nines = map[string]bool{}
		uphill(zero[0], zero[1])
		total += len(nines)
	}

	fmt.Println(total)
}
