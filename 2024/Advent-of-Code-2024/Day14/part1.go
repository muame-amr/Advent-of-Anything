package main

import "fmt"

func solveA(positions [][2]int, velocity [][2]int, w int, h int) {
	c := make(map[[2]int]int)
	n := len(positions)
	mid_x := w / 2
	mid_y := h / 2

	for i := 0; i < n; i++ {
		x_final := (positions[i][0] + velocity[i][0]*100) % w
		y_final := (positions[i][1] + velocity[i][1]*100) % h

		if x_final < 0 {
			x_final += w
		}

		if y_final < 0 {
			y_final += h
		}

		if x_final != mid_x && y_final != mid_y {
			c[[2]int{x_final, y_final}]++
		}

	}

	quad := make([]int, 4)
	for k, v := range c {
		if k[0] < mid_x && k[1] < mid_y {
			quad[0] += v
		} else if k[0] > mid_x && k[1] < mid_y {
			quad[1] += v
		} else if k[0] < mid_x && k[1] > mid_y {
			quad[2] += v
		} else if k[0] > mid_x && k[1] > mid_y {
			quad[3] += v
		}
	}
	fmt.Println(quad[0] * quad[1] * quad[2] * quad[3])
}
