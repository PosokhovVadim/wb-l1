package main

import "fmt"

func main() {
	var num int64
	var i int
	fmt.Scanf("%d", &num)
	fmt.Scanf("%d", &i)

	res := num | (int64(1) << i)
	fmt.Println(res)
}
