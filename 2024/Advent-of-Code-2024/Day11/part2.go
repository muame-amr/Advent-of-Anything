package main

import (
	"fmt"
	"strconv"
)

func solveB(arrangement []string) {
	stones := map[string]int{}
	for _, stonestr := range arrangement {
		stones[stonestr] = 1
	}

	for i := 0; i < 75; i++ {
		newstones := map[string]int{}
		for stone, count := range stones {
			if stone == "0" {
				newstones["1"] += count
			} else if len(stone)%2 == 0 {
				var midpt int = len(stone) / 2
				l, _ := strconv.Atoi(stone[:midpt])
				r, _ := strconv.Atoi(stone[midpt:])
				newstones[fmt.Sprintf("%d", l)] += count
				newstones[fmt.Sprintf("%d", r)] += count
			} else {
				stoneInt, _ := strconv.Atoi(stone)
				newstones[fmt.Sprintf("%d", stoneInt*2024)] += count
			}
		}
		stones = newstones
	}

	total := 0
	for _, val := range stones {
		total += val
	}
	fmt.Println(total)
}
