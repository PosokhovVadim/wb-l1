package main

import (
	"fmt"
	"strings"
)

func uniqueLetters(str string) bool {
	var set = make(map[string]bool)
	for i := 0; i < len(str); i++ {
		v := strings.ToLower(string(str[i]))

		if set[v] {
			return false
		}
		set[v] = true
	}

	return true
}

func main() {
	str := "abCdefAaf"

	fmt.Println(uniqueLetters(str))
}
