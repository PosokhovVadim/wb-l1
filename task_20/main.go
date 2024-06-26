package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	slices.Reverse(words)

	return strings.Join(words, " ")
}

func main() {

	fmt.Println(reverseWords("snow dog sun"))
}
