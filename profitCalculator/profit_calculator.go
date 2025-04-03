package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	
	printMessage("Profit Calculator")

	revenue , err := getUserInput("Enter Revenue:")

	if err != nil{
		printError(err)
		return
	}
	expenses , err := getUserInput("Enter Expenses:")

	if err != nil{
		printError(err)
		return
	}
	taxRate , err := getUserInput("Enter Expected Tax Rate:")
	if err != nil{
		printError(err)
		return
	}

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
	addResultToFile(ebt, profit, ratio)
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

func getUserInput(promptText string) (float64, error){
	var userInput float64
	fmt.Println(promptText)
	fmt.Scan(&userInput)

	if userInput <= 0{
		return 0, errors.New("invalid input. amount cannot be smaller than zero or zero")
	}
	return userInput , nil
}

func addResultToFile(ebt float64, profit float64, ratio float64){
	values := fmt.Sprint("EBT: ",ebt,"\nProfit: ", profit,"\nRatio: ",ratio)
	os.WriteFile(("profit.txt"), []byte(values),0644)
}

func printError(message error){
	fmt.Println(message)
}

