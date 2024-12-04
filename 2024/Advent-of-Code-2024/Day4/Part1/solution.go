package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("input.in")
	defer file.Close()

	var str string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text() + "\n"
	}

	lines := strings.Split(str, "\n")
	grid := make([][]string, len(lines))

	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	total := 0
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == "X" {
				total += findXmasHorizontal(grid, j, i)
				total += findXmasHorizontalReverse(grid, j, i)
				total += findXmasVertical(grid, j, i)
				total += findXmasVerticalReverse(grid, j, i)
				total += findXmasNE(grid, j, i)
				total += findXmasSE(grid, j, i)
				total += findXmasNW(grid, j, i)
				total += findXmasSW(grid, j, i)
			}
		}
	}

	fmt.Println(total)
}

func findXmasHorizontal(grid [][]string, x int, y int) int {
	n := len(grid[0])

	if x > n-4 {
		return 0
	}

	if grid[y][x+1] == "M" && grid[y][x+2] == "A" && grid[y][x+3] == "S" {
		return 1
	}

	return 0
}

func findXmasHorizontalReverse(grid [][]string, x int, y int) int {
	if x < 3 {
		return 0
	}

	if grid[y][x-1] == "M" && grid[y][x-2] == "A" && grid[y][x-3] == "S" {
		return 1
	}

	return 0
}

func findXmasVertical(grid [][]string, x int, y int) int {
	n := len(grid)
	if y > n-4 {
		return 0
	}

	if grid[y+1][x] == "M" && grid[y+2][x] == "A" && grid[y+3][x] == "S" {
		return 1
	}

	return 0
}

func findXmasVerticalReverse(grid [][]string, x int, y int) int {
	if y < 3 {
		return 0
	}

	if grid[y-1][x] == "M" && grid[y-2][x] == "A" && grid[y-3][x] == "S" {
		return 1
	}

	return 0
}

func findXmasSE(grid [][]string, x int, y int) int {
	n := len(grid)
	m := len(grid[0])

	if x > n-4 || y > m-4 {
		return 0
	}

	if grid[y+1][x+1] == "M" && grid[y+2][x+2] == "A" && grid[y+3][x+3] == "S" {
		return 1
	}

	return 0
}

func findXmasNE(grid [][]string, x int, y int) int {
	n := len(grid)

	if x > n-4 || y < 3 {
		return 0
	}

	if grid[y-1][x+1] == "M" && grid[y-2][x+2] == "A" && grid[y-3][x+3] == "S" {
		return 1
	}

	return 0
}

func findXmasSW(grid [][]string, x int, y int) int {
	m := len(grid[0])

	if x < 3 || y > m-4 {
		return 0
	}

	if grid[y+1][x-1] == "M" && grid[y+2][x-2] == "A" && grid[y+3][x-3] == "S" {
		return 1
	}

	return 0
}

func findXmasNW(grid [][]string, x int, y int) int {
	if x < 3 || y < 3 {
		return 0
	}

	if grid[y-1][x-1] == "M" && grid[y-2][x-2] == "A" && grid[y-3][x-3] == "S" {
		return 1
	}

	return 0
}
