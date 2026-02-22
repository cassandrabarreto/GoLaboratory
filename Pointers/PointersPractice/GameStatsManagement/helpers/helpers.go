package helpers

import (
	"errors"
	"slices"
)


func LookForInventoryElementIndex(inventory []string, element string, index int) (int, error){
	if index >= len(inventory) { 
        return -1, nil
    }
	if !slices.Contains(inventory, element){
		return -2, errors.New("Element not in inventory.")
	}

	if inventory[index] == element {
		return index, nil
	}

	return LookForInventoryElementIndex(inventory, element, index+1)
}

