package main

import "fmt"

func solveA(grid [][]rune) {
	total := 0
	directons := [][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	visited := make(map[[2]int]bool)
	var expand func(x int, y int, plant rune) int
	expand = func(x int, y int, plant rune) int {
		area := 0
		perimeter := 0
		queue := [][2]int{{x, y}}
		visited[[2]int{x, y}] = true

		for len(queue) > 0 {
			pop := queue[0]
			queue = queue[1:]
			area++

			for _, dir := range directons {
				dx := pop[0] + dir[0]
				dy := pop[1] + dir[1]

				// for each blocked direction we add fence/perimeter
				if dx < 0 || dy < 0 || dx >= len(grid[0]) || dy >= len(grid) || grid[dy][dx] != plant {
					perimeter++
					continue
				}

				// add to queue for each direction for not visited vertex
				if !visited[[2]int{dx, dy}] {
					queue = append(queue, [2]int{dx, dy})
					visited[[2]int{dx, dy}] = true
				}
			}
		}
		return area * perimeter
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[[2]int{j, i}] {
				total += expand(j, i, grid[i][j])
			}
		}
	}

	fmt.Println(total)
}

func display(grid [][]rune) {
	for _, line := range grid {
		for _, plant := range line {
			fmt.Printf("%c", plant)
		}
		fmt.Println()
	}
}
