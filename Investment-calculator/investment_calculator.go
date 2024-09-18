package main

import "math"
import "fmt"

func main() {
	var investmentAmount, years = 1000, 10
	expectedReturnRate := 5.5

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Print(futureValue)

}
