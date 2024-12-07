package main

import (
	"fmt"
	"os"
	"strings"
)

var coords = make(map[string]bool)
var total int = 0

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func main() {
	content, _ := os.ReadFile("test.in")
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
	fmt.Println("total: ", move(grids, startX, startY, currentDir))
}

func move(grids [][]string, x int, y int, currentDir Direction) int {
	coords[fmt.Sprintf("%d,%d", x, y)] = true
	grids[y][x] = "X"

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
		return len(coords)
	}

	if nextX >= 0 && nextX < len(grids[0]) && nextY >= 0 && nextY < len(grids) {
		if grids[nextY][nextX] == "#" {
			// Rotate 90 degrees clockwise
			currentDir = (currentDir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}

	return move(grids, x, y, currentDir)
}
