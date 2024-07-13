package main

import "fmt"



func main() {
    var taxRate float64
	var expenses float64
    var revenue float64

	
	taxRate = getUserInput("Tax rate: ")
	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	

	beta,beta1 := returnBeta(revenue,taxRate)

	fmt.Printf("Beta: %.1f\n",beta)
	fmt.Printf("Beta1: %.1f",beta1)
	fmt.Printf("expenses: %.1f",expenses)

}

func getUserInput(infoText string) float64 {
	var value float64
	fmt.Print(infoText)
	fmt.Scan(&value)

	return value
}



func returnBeta(revenue float64,taxRate float64)( float64, float64) {
	fv :=  revenue / taxRate
	fvt :=  revenue * taxRate

	return fv,fvt
}