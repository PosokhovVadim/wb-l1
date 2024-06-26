package main

import (
	"fmt"
	"reflect"
)

func solution1(data interface{}) {
	fmt.Printf("%T\n", data)

}

func solution2(data interface{}) {
	t := reflect.TypeOf(data)

	fmt.Println(t)
}

func solution3(data interface{}) {
	switch t := data.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		_ = t
		fmt.Println("unknown")
	}
}
func main() {
	var k int
	fmt.Scanf("%d", &k)
	// str := "qqq"
	// num := 1
	ch := make(chan int)
	switch k {
	case 1:
		solution1(ch)
	case 2:
		solution2(ch)
	case 3:
		solution3(ch)
	default:
		solution1(ch)
	}
}
