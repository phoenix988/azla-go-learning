package main

import (
	"fmt"
	"strconv"
)

func convertToNum(word string) int {
	// Convert selected word count to int
	convertToNum, Err := strconv.Atoi(word)

	if Err != nil {
		fmt.Println("Error:", Err)
	} else {
		return convertToNum
	}

	return 0
}