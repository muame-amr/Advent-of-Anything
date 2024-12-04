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
	for i := 1; i < len(grid)-2; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == "A" {
				total += findXMAS(grid, j, i)
			}
		}
	}

	fmt.Println(total)
}

func findXMAS(grid [][]string, x int, y int) int {
	if grid[y-1][x-1] == "M" && grid[y+1][x+1] == "S" {
		if grid[y-1][x+1] == "M" && grid[y+1][x-1] == "S" {
			return 1
		}
		if grid[y-1][x+1] == "S" && grid[y+1][x-1] == "M" {
			return 1
		}
	}
	if grid[y-1][x-1] == "S" && grid[y+1][x+1] == "M" {
		if grid[y-1][x+1] == "M" && grid[y+1][x-1] == "S" {
			return 1
		}
		if grid[y-1][x+1] == "S" && grid[y+1][x-1] == "M" {
			return 1
		}
	}
	return 0
}
