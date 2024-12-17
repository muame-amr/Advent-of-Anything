package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	content, _ := os.ReadFile("input.in")
	input := strings.Split(string(content), "\r\n\r\n")
	re := regexp.MustCompile(`(\d+)`)
	machines := [][]float64{}
	for _, parts := range input {
		coordstr := re.FindAllString(parts, -1)
		coordint := []float64{}
		for _, coord := range coordstr {
			num, _ := strconv.ParseFloat(coord, 64)
			coordint = append(coordint, num)
		}
		machines = append(machines, coordint)
	}
	// solveA(machines)
	solveB(machines)
	elapsed := time.Since(start)
	fmt.Println("Execution time : ", elapsed)
}
