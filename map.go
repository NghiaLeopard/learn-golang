package main

import "fmt"

type transform func(int) int

func main() {
	emptyArr := []int{}
	numbers := []int{1,2,3,4,5}

	newArr := append(emptyArr,numbers...)



	result := sumup(1,2,3,4,5)

	fmt.Println(newArr)
	fmt.Println(result)
}

func sumup(numbers ...int) int {
	fmt.Print(numbers)
	return 0
}
