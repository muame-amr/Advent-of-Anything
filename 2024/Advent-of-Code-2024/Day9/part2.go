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
	blocks := make(map[int]string)
	fileId := 0
	idx := 0
	fileBlock := make(map[int]tuple)
	freeBlock := []tuple{}

	for pos, char := range strings.Split(fileContent, "") {
		diskMap, _ := strconv.Atoi(char)
		if pos%2 == 0 {
			st := idx
			for i := 0; i < diskMap; i++ {
				blocks[idx] = fmt.Sprintf("%d", fileId)
				idx++
			}
			ed := idx - 1
			fileBlock[fileId] = tuple{start: st, end: ed}
			fileId++
		} else {
			st := idx
			for i := 0; i < diskMap; i++ {
				blocks[idx] = "."
				idx++
			}
			ed := idx - 1
			freeBlock = append(freeBlock, tuple{start: st, end: ed})
		}
	}

	displayDiskMap(blocks, idx)
	// fmt.Println(fileBlock)
	// fmt.Println(freeBlock)

	for curr := fileId - 1; curr >= 0; curr-- {
		fileSize := fileBlock[curr].end - fileBlock[curr].start + 1
		for _, dot := range freeBlock {
			freeSize := dot.end - dot.start + 1
			println(curr, fileSize, freeSize)
			if freeSize > 0 && fileSize <= freeSize {
				sdx := dot.start
				edx := fileBlock[curr].end
				for x := 0; x < fileSize; x++ {
					blocks[sdx] = fmt.Sprintf("%d", curr)
					sdx++
					blocks[edx] = "."
					edx--
				}
				dot.start = sdx
			}
		}
	}

	displayDiskMap(blocks, idx)

	// total := 0
	// blockId := 0
	// for blocks[blockId] != "." {
	// 	val, _ := strconv.Atoi(blocks[blockId])
	// 	total += blockId * val
	// 	blockId++
	// }
	// fmt.Println(total)
}
