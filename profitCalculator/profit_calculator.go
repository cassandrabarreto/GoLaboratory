package main

import (
	"fmt"
)

func main() {
	
	printMessage("Profit Calculator")

	revenue := getUserInput("Enter Revenue:")
	expenses := getUserInput("Enter Expenses:")
	taxRate := getUserInput("Enter Expected Tax Rate:")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Print("EBT:", ebt, "\n")
	fmt.Print("Profit:", profit, "\n")
	fmt.Printf("%.1f\n", ratio)
}

func printMessage(message string){
	fmt.Println(message)
}

func calculateFinancials(revenue float64 , expenses float64, taxRate float64)(float64, float64, float64){
	ebt := calculateEbt(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	return ebt, profit, ratio

}

func calculateProfit(ebt float64, taxRate float64) float64{
	result := ebt * (1 - taxRate/100)
	return result
}

func calculateEbt(revenue float64, expenses float64) float64{
	result := revenue - expenses
	return result
}

func calculateRatio(ebt float64, profit float64) float64{
	result := ebt / profit
	return result
}

func getUserInput(promptText string) float64 {
	var userInput float64
	fmt.Println(promptText)
	fmt.Scan(&userInput)
	return userInput
}

