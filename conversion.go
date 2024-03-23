package main

import (
	"fmt"
	"strconv"
	"math/rand"
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


func generateSessionID() string {
    // Generate a random session ID (dummy example)
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, 32)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


