package main

import "fmt"


func profitCalculator(){

var revenue float64
var expenses float64
var taxRate float64
var earningsBeforeTax float64
var earningsAfterTax float64
var ratio float64


fmt.Print("What is the revenue: ")
fmt.Scan(&revenue)
fmt.Print("What is the Expenses: ")
fmt.Scan(&expenses)
fmt.Print("What is the Tax Rate: ")
fmt.Scan(&taxRate)


earningsBeforeTax = revenue - expenses
earningsAfterTax = earningsBeforeTax - (earningsBeforeTax*(taxRate/100))
ratio = earningsBeforeTax/earningsAfterTax
fmt.Print("You Earnings Before Tax are: ")
fmt.Println(earningsBeforeTax)

fmt.Print("Your Earnings After Tax are: ")
fmt.Println(earningsAfterTax)

fmt.Print("Your Tax Ratio is: ")
fmt.Println(ratio)

}