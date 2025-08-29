package to_do

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ToDo struct{
	Text string `json:"text"`
}

// Create constructor
func NewToDo(content string) (ToDo, error){

	if content == ""{
		//Return dummy Note struct.
		return ToDo{}, errors.New("invalid input")
	}
	return ToDo{
		Text : content,
	}, nil
}

// Create 'class' (struct ) method
func (todo ToDo) Display() {
	fmt.Println(todo.Text)
}

func (todo ToDo)Save() error{
	filename := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil{
		return err
	}
	return os.WriteFile(filename, json, 0644)

}