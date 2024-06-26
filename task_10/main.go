package main

import (
	"fmt"
	"math"
)

func getGroup(temp float64) int64 {
	return int64(math.Round(temp)/10) * 10
}
func main() {

	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	data := make(map[int64][]float64)
	for i := 0; i < len(temp); i++ {
		group := getGroup(temp[i])
		if _, ok := data[group]; !ok {
			data[group] = make([]float64, 0)
		}
		data[group] = append(data[group], temp[i])
	}

	fmt.Println(data)
}
