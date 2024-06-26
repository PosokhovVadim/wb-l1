package main

import (
	"fmt"
	"unsafe"
)

var justString string

func createHugeString(n int) []byte {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = 'v'
	}

	return s
}

func someFunc() {
	v := createHugeString(1 << 10)

	justString = string(v[:100])
	v = nil
}

func main() {
	someFunc()
	fmt.Println(justString)
	fmt.Println(unsafe.Pointer(&justString))
}
