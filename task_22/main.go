package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(4294967296)
	b := big.NewInt(4194304)

	mul := new(big.Int)
	div := new(big.Int)
	sum := new(big.Int)
	sub := new(big.Int)

	fmt.Println("multyplication: ", mul.Mul(a, b))
	fmt.Println("division: ", div.Div(a, b))
	fmt.Println("sum: ", sum.Add(a, b))
	fmt.Println("sub: ", sub.Sub(a, b))
}
