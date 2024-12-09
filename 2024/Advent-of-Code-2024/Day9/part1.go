package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveA(fileContent string) {
	blocks := []rune{}
	fileId := 0
	idx := 0
	freeQueue := []int{}
	fragStack := []int{}

	for pos, char := range strings.Split(fileContent, "") {
		diskMap, _ := strconv.Atoi(char)
		if pos%2 == 0 {
			for i := 0; i < diskMap; i++ {
				blocks = append(blocks, rune('0'+fileId))
				fragStack = append(fragStack, idx)
				idx++
			}
			fileId++
		} else {
			for i := 0; i < diskMap; i++ {
				blocks = append(blocks, rune('.'))
				freeQueue = append(freeQueue, idx)
				idx++
			}
		}
	}

	lastIdx := len(fragStack) - 1
	for freeQueue[0] <= fragStack[lastIdx] {
		blocks[freeQueue[0]] = blocks[fragStack[lastIdx]] // swap
		blocks[fragStack[lastIdx]] = '.'
		freeQueue = freeQueue[1:] // dequeue
		lastIdx--                 // pop stack
	}

	total := 0
	blockId := 0
	for blocks[blockId] != '.' {
		total += blockId * int(blocks[blockId]-'0')
		blockId++
	}
	fmt.Println(total)
}
