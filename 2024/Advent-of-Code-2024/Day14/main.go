package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	width := 101
	height := 103
	content, _ := os.ReadFile("input.in")
	input := strings.Split(string(content), "\r\n")
	positions := [][2]int{}
	velocity := [][2]int{}
	for _, line := range input {
		st := strings.Split(line, " ")[0]
		nd := strings.Split(line, " ")[1]
		pos := strings.Split(st, "=")[1]
		vel := strings.Split(nd, "=")[1]
		p := strings.Split(pos, ",")
		v := strings.Split(vel, ",")
		xp, _ := strconv.Atoi(p[0])
		yp, _ := strconv.Atoi(p[1])
		xv, _ := strconv.Atoi(v[0])
		yv, _ := strconv.Atoi(v[1])
		positions = append(positions, [2]int{xp, yp})
		velocity = append(velocity, [2]int{xv, yv})
	}
	// solveA(positions, velocity, width, height)
	solveB(positions, velocity, width, height)
	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}
