package filemanager

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetBalanceFromFile(filename string) (float64, error){
	data , err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("failed to find balance")
		//panic("Exit the program.")
	} 
	balanceText := string(data)
	balance , _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}

func WriteBalanceToFile(balance float64, fileName string){
	balanceText := fmt.Sprint(balance)
	os.WriteFile((fileName), []byte(balanceText), 0644)
}