package main

import "math"
import "fmt"

func main() {

	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Println("Future value: ",futureValue)
	//fmt.Println("Real value: ",futureRealValue)
	fmt.Printf("Future value: %v\nFuture Value (adjusted for Inflation): %v", futureValue, futureRealValue)

}
