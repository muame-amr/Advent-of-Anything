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

	for curr := fileId - 1; curr >= 0; curr-- {
		fileSize := fileBlock[curr].end - fileBlock[curr].start + 1
		for i, dot := range freeBlock {
			freeSize := dot.end - dot.start + 1
			if freeSize > 0 && fileSize <= freeSize {
				for x := 0; x < fileSize; x++ {
					blocks[dot.start+x] = fmt.Sprintf("%d", curr)
				}
				for x := 0; x < fileSize; x++ {
					blocks[fileBlock[curr].end-x] = "."
				}
				freeBlock[i] = tuple{start: dot.start + fileSize, end: dot.end}
				break
			}
		}
	}

	checksum := 0
	for k, v := range blocks {
		if v != "." {
			f, _ := strconv.Atoi(v)
			checksum += k * f
		}
	}
	fmt.Println(checksum)
}

// 8553289251139 : too high
