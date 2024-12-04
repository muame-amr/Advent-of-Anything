package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("test.in")
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
		for j := 1; j < len(grid[0])-2; j++ {
			if grid[i][j] == "A" {
				println(grid[i-1][j-1], "", grid[i-1][j+1])
				println("", grid[i][j], "")
				println(grid[i+1][j-1], "", grid[i+1][j+1])
				total += findXMAS(grid, j, i)
			}
		}
	}

	fmt.Println(total)
}

func findXMAS(grid [][]string, x int, y int) int {
	n := len(grid)
	m := len(grid[0])

	if x > n-2 || y > m-2 || x < 1 || y < 1 {
		return 0
	}

	if (grid[y-1][x-1] == "M" && grid[y+1][x+1] == "S") ||
		(grid[y-1][x-1] == "S" && grid[y+1][x+1] == "M") ||
		(grid[y+1][x-1] == "S" && grid[y-1][x+1] == "M") ||
		(grid[y+1][x-1] == "M" && grid[y-1][x+1] == "S") {
		return 1
	}

	return 0
}
