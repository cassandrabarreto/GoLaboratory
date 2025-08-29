package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/to_do"
)

func main () {
	title, content := getNoteData()
	todoText := getUserInput("To do Text")
	text := getUserInput("Type something")
	

	todo, err := to_do.NewToDo(todoText)
	
	if err != nil {
    		fmt.Println("Error creating ToDo:", err)
    	return
	}
	
	todo.Display()

	err = todo.Save()

	if err != nil {
    		fmt.Println("Error saving ToDo list.")
    	return
	}

	userNote, err := note.NewNote(title, content)
	if err != nil {
    		fmt.Println("Error creating Note:", err)
    	return
	}

	userNote.Display()

	err = userNote.Save()
	
	if err != nil {
    		fmt.Println("Error saving Note.")
    	return
	}

	fmt.Println("Note Saved successfully.")
	
}

/* Function that retrieves user input.*/
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	
	/* \n We cannot use fmt.Scan because Scan can only read 1 value.*/
	reader := bufio.NewReader(os.Stdin)

	/* \n tells ReadString stop reading when there is a line break*/
	text, err := reader.ReadString('\n')

	if err != nil{
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}


func getTodoData() string{
	return getUserInput("Todo text:")
}


/* Function that returns two parts: title and Note.*/
func getNoteData()(string, string){
	title := getUserInput("Note Title:")

	content := getUserInput("Content:")

	return title, content
}