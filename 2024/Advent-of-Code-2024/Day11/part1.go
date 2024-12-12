package main

import (
	"fmt"
	"strconv"
)

func solveA(arrangement []string) {
	stones := []int{}
	for _, stonestr := range arrangement {
		stone, _ := strconv.Atoi(stonestr)
		stones = append(stones, stone)
	}

	for i := 0; i < 75; i++ {
		newstones := []int{}
		for _, stone := range stones {
			if stone == 0 {
				newstones = append(newstones, 1)
			} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
				stonestr := fmt.Sprintf("%d", stone)
				var midpt int = len(stonestr) / 2
				l, _ := strconv.Atoi(stonestr[:midpt])
				r, _ := strconv.Atoi(stonestr[midpt:])
				newstones = append(newstones, l, r)
			} else {
				newstones = append(newstones, stone*2024)
			}
		}
		stones = newstones
	}

	fmt.Println(len(stones))
}
