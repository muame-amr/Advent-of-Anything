package main

import (
	"fmt"
)

type antenna struct {
	char rune
	x    int
	y    int
}

func solveA(fileContent []string) {
	antennas := []antenna{}
	y_size := len(fileContent)
	x_size := len(fileContent[0])

	for y, line := range fileContent {
		for x, char := range line {
			if char != '.' {
				antennas = append(antennas, antenna{char, x, y})
			}
		}
	}

	antinodes := map[string]bool{}
	for _, a := range antennas {
		for _, b := range antennas {
			if a != b && a.char == b.char {
				dy := b.y - a.y
				dx := b.x - a.x
				nodeY := b.y + dy
				nodeX := b.x + dx

				if nodeX >= 0 &&
					nodeX < x_size &&
					nodeY >= 0 &&
					nodeY < y_size {
					antinodes[fmt.Sprintf("%d,%d", nodeX, nodeY)] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
