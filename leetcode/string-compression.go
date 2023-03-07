package leetcode

import (
	"fmt"
	"strconv"
)

func Compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	currentLength := 1
	currentChar := chars[0]
	fmt.Printf("%c\n", currentChar)
	currentIdx := 0

	for i := 1; i <= len(chars); i++ {
		if i == len(chars) || chars[i] != currentChar {
			update(chars, &currentIdx, currentChar, currentLength)

			if i < len(chars) {
				currentLength = 1
				currentChar = chars[i]
			}
		} else {
			currentLength += 1
		}
	}
	chars = chars[:currentIdx]
	return currentIdx
}

func update(chars []byte, idx *int, c byte, numC int) {
	chars[*idx] = c
	*idx++

	if numC == 1 {
		return
	}

	numCAsBytes := []byte(strconv.Itoa(numC))
	for pos, val := range numCAsBytes {
		chars[*idx+pos] = val
	}
	*idx += len(numCAsBytes)
}
