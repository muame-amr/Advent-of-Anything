package main

import (
	"fmt"
	"math"
)

func solveB(machines [][]float64) {
	var calcTokens func(coords []float64) int
	calcTokens = func(coords []float64) int {
		var aX float64 = coords[0]
		var aY float64 = coords[1]
		var bX float64 = coords[2]
		var bY float64 = coords[3]
		var pX float64 = coords[4]
		var pY float64 = coords[5]

		var a float64 = (pX*bY - pY*bX) / (aX*bY - aY*bX)
		var b float64 = (aX*pY - aY*pX) / (aX*bY - aY*bX)

		if math.Floor(a) == a && math.Floor(b) == b {
			return int(a)*3 + int(b)
		} else {
			return 0
		}
	}

	total := 0
	for _, machine := range machines {
		machine[4] += math.Pow(10, 13)
		machine[5] += math.Pow(10, 13)
		total += calcTokens(machine)
	}
	fmt.Println(total)
}
