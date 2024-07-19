package main

import (
	// "errors"
	// "fmt"
	// "os"
	// "strconv"

	// "errors"
	// "os"
	// "strconv"

	"example.com/main/fileops"
)

// func getBalance()(float64,error) {

// 	data, err :=os.ReadFile("balance.txt")

// 	if err != nil {
// 		return 1000,errors.New("Have not enough money")
// 	}

// 	dataParse := string(data)

// 	balanceText,_ := strconv.ParseFloat(dataParse,64)

// 	if err != nil {
// 		return balanceText,errors.New("Failed to parse store value")
// 	}

// 	return balanceText,nil
// }

// func getUserInput(textField string)(float64,error) {
// 	var value float64
// 	fmt.Print(textField)
// 	fmt.Scan(&value)

// 	if value < 0 {
// 		return value,errors.New("Value is invalid")
// 	}

// 	return value,nil
// }

// func test1(value,value1 float64, infoText string) {

// }


func main() {
	// // var number int = 0


	// var number int

	// fmt.Print("Please enter number: ")
	// fmt.Scan(&number)

	//  balance := fmt.Sprint(number)

	// os.WriteFile("balance.txt",[]byte(balance),0644)
	// // for {
	// // 	fmt.Printf("Number: %.d\n", number)

	// // 	if(number >= 0) {
	// // 	   number++
	// // 		fmt.Printf("Number: %.d\n", number)
	// // 		return 
	// // 	}
	// // }

	// balanceText,err := getBalance()

	// if err != nil {
	// fmt.Println(err)

	// }

	// fmt.Println(balanceText)

	fileops.PrintText()

}