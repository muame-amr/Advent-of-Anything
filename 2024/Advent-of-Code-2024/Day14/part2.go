package main

import "fmt"

func solveB(positions [][2]int, velocity [][2]int, w int, h int) {
	n := len(positions)

	for l := 1; l < 100_000; l++ {
		fmt.Println(l)
		c := make(map[[2]int]bool)
		grid := createGrid(w, h)
		dupe := false

		for i := 0; i < n; i++ {
			positions[i][0] = (positions[i][0] + velocity[i][0]) % w
			positions[i][1] = (positions[i][1] + velocity[i][1]) % h

			for positions[i][0] < 0 {
				positions[i][0] += w
			}

			for positions[i][1] < 0 {
				positions[i][1] += h
			}

			if c[[2]int{positions[i][0], positions[i][1]}] {
				dupe = true
				continue
			}

			c[[2]int{positions[i][0], positions[i][1]}] = true
			grid[positions[i][1]][positions[i][0]] = "#"
		}

		if dupe {
			continue
		} else {
			displayGrid(grid)
			break
		}
	}
}

func createGrid(w int, h int) [][]string {
	grid := make([][]string, h)
	for i := 0; i < h; i++ {
		row := make([]string, w)
		for j := 0; j < w; j++ {
			row[j] = "."
		}
		grid[i] = row
	}

	return grid
}

func displayGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
