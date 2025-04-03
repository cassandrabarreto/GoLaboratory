package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
const accountBalance = "balance.txt"

func getBalanceFromFile() (float64, error){
	data , err := os.ReadFile(accountBalance)

	if err != nil {
		return 1000, errors.New("failed to find balance")
		//panic("Exit the program.")
	} 
	balanceText := string(data)
	balance , _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}

func main(){

	fmt.Print("Bank App\n")
	
	var totalBalance , err = getBalanceFromFile()

	if err != nil{
		fmt.Println("ERROR")
	}
	
	for { 
		showMenu()
		choice := getUserInput("Enter your choice:")
		
		switch choice {
		case 1:
			var depositAmount float64
			fmt.Print("Enter amount to deposit:\n")
			fmt.Scan(&depositAmount)
			
			if depositAmount < 0 {
				fmt.Print("\nInvalid Amount. Please try again.\n\n")
				continue
			}
			totalBalance += depositAmount
			writeBalanceToFile(totalBalance, accountBalance)
			fmt.Print("Total Balance:", totalBalance, "\n")	
		case 2:
			var withdrawAmount float64
			fmt.Print("Enter a mount to withdraw:\n")
			fmt.Scan(&withdrawAmount)
			
			if withdrawAmount > totalBalance {
				fmt.Print("\nInvalid Amount. Please try again.\n\n")
				continue
			}
			totalBalance -= withdrawAmount
			fmt.Print("Total Balance:", totalBalance, "\n")
		case 3:
			fmt.Print("Your balance:\n")
			fmt.Print("Total Balance:", totalBalance, "\n")
		default:
			fmt.Print("Exiting...\n")
			fmt.Print("Thanks for choosing our bank!\n")
			return	
		}
	}
}

func getUserInput(promptText string) float64 {
	var userInput float64
	fmt.Println(promptText)
	fmt.Scan(&userInput)
	return userInput
}

func writeBalanceToFile(balance float64, fileName string){
	balanceText := fmt.Sprint(balance)
	os.WriteFile((fileName), []byte(balanceText), 0644)
}


