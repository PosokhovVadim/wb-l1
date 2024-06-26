package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make([]string, 0)
	m := make(map[string]bool)

	for _, v := range strings {
		m[v] = true
	}

	for k := range m {
		set = append(set, k)
	}

	fmt.Println(set)
}
