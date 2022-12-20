package main

import (
	"fmt"
	"test-ut/internal/calculate/substract"
	"test-ut/internal/calculate/sum"
)

func main() {
	result := sum.Add(1, 2)
	fmt.Println(result)
	result1 := substract.Substract(3, 2)
	fmt.Println(result1)
}
