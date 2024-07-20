package main

import "fmt"

func main() {
	user := 23

	var ptn *int = &user

	fmt.Println("Pointer:",*ptn)

	calculate(ptn)

	fmt.Println("Age: ",user)

}

// input: pointer to integer
func calculate(age *int) {
	*age = *age - 12
}