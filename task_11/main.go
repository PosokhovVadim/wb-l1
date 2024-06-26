package main

import "fmt"

func arrToMap(arr []int64) map[int64]int64 {
	m := make(map[int64]int64)
	for _, v := range arr {
		m[v] = v
	}
	return m
}

func main() {
	X := []int64{1, 6, 7, 8, 11, 15, 18}
	Y := []int64{1, 2, 9, 10, 11, 12}
	XMap := arrToMap(X)
	res := make([]int64, 0)
	for _, v := range Y {
		if _, ok := XMap[v]; ok {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}
