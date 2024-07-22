package main

import "fmt"

func main() {
	value := add(1, 2)

	fmt.Println(value)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}