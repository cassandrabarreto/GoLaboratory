package main

import (
	"fmt"
	"strings"
)


func lengthOfLastWord(s string) int {
	phrase := strings.TrimSpace(s)
	wordList := strings.Split(phrase, " ")
	if len(wordList) == 0 {
		return 0
	}
	return len(wordList[len(wordList)-1])
}


func checkWord(wordList []string, index int) int {
	if index >= len(wordList) {
		return len(wordList[index-1])
	}
	return checkWord(wordList, index+1)
}

func main() {
	testCases := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
		"",
		"single",
	}

	for _, s := range testCases {
		fmt.Printf("Input: %q, Length of Last Word: %d\n", s, lengthOfLastWord(s))
	}
}
