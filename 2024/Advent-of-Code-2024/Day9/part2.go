package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tuple struct {
	start int
	end   int
}

func solveB(fileContent string) {
	diskmap := []int{}
	blocks := [][]string{}
	checksum := 0

	for _, val := range strings.Split(fileContent, "") {
		num, _ := strconv.Atoi(val)
		diskmap = append(diskmap, num)
	}

	// fmt.Println(diskmap)

	for i, inp := range diskmap {
		file := []string{}
		for j := inp; j > 0; j-- {
			if i%2 == 0 {
				file = append(file, fmt.Sprintf("%d", i/2))
			} else {
				file = append(file, ".")
			}
		}
		if len(file) > 0 {
			blocks = append(blocks, file)
		}
	}

	// fmt.Println(blocks)

	movedId := map[string]bool{}

outer:
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i][0] != "." && movedId[blocks[i][0]] == false {
			for j := 0; j <= i; j++ {
				if blocks[j][0] == "." {
					movedId[blocks[i][0]] = true
					if len(blocks[j]) == len(blocks[i]) {
						temp := make([]string, len(blocks[j]))
						copy(temp, blocks[j])
						blocks[j] = blocks[i]
						blocks[i] = temp
						continue outer
					} else if len(blocks[j]) > len(blocks[i]) {
						temp := make([]string, len(blocks[i]))
						copy(temp, blocks[i])
						for k := range blocks[i] {
							blocks[i][k] = "."
						}
						blocks = append(blocks[:j], append([][]string{temp, blocks[j][len(blocks[i]):]}, blocks[j+1:]...)...)
						// fmt.Println("gt", i, blocks)
						i++
						continue outer
					}
				}
			}
		}
	}
	// fmt.Println(blocks)

	fragmented := []string{}
	for _, block := range blocks {
		fragmented = append(fragmented, block...)
	}

	// fmt.Println(fragmented)

	for i, file := range fragmented {
		if file != "." {
			num, _ := strconv.Atoi(file)
			checksum += i * num
		}
	}

	fmt.Println(checksum)
}

// 6398096697992
// 8553289251139 : too high
