package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func main() {
	start := time.Now()
	content, _ := os.ReadFile("input.in")
	fileContent := string(content)

	grids := [][]string{}
	startX, startY := 0, 0

	for i, line := range strings.Fields(fileContent) {
		pixels := []string{}
		for j, cell := range strings.Split(line, "") {
			if cell == "^" {
				startX = j
				startY = i
			}
			pixels = append(pixels, cell)
		}
		grids = append(grids, pixels)
	}

	var currentDir Direction = Up
	total := 0
	for i := 0; i < len(grids); i++ {
		for j := 0; j < len(grids[0]); j++ {
			coords := map[string]bool{}
			if grids[i][j] == "." {
				grids[i][j] = "#"
				if move(grids, startX, startY, currentDir, coords) {
					total++
				}
				grids[i][j] = "."
			}
		}
	}
	fmt.Println(total)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func move(grids [][]string, x int, y int, currentDir Direction, coords map[string]bool) bool {
	coords[fmt.Sprintf("%d,%d,%v", x, y, currentDir)] = true

	nextX, nextY := x, y
	switch currentDir {
	case Up:
		nextY--
	case Right:
		nextX++
	case Down:
		nextY++
	case Left:
		nextX--
	}

	if nextX < 0 || nextX >= len(grids[0]) || nextY < 0 || nextY >= len(grids) {
		return false
	}

	if nextX >= 0 && nextX < len(grids[0]) && nextY >= 0 && nextY < len(grids) {
		if grids[nextY][nextX] == "#" {
			// Rotate 90 degrees clockwise
			currentDir = (currentDir + 1) % 4
		} else {
			x, y = nextX, nextY
			if coords[fmt.Sprintf("%d,%d,%v", x, y, currentDir)] {
				return true
			}
		}
	}

	return move(grids, x, y, currentDir, coords)
}

func display(grids [][]string) {
	for _, line := range grids {
		for _, cell := range line {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
