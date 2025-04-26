package main

import (
	"fmt"

	"example.com/note/note"
)

func main () {
	title, content  := getNoteData()
	userNote , err := note.NewNote(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	
}

/* Function that retrieves user input.*/
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)
	return value
}

/* Function that returns two parts: title and Note.*/
func getNoteData()(string, string){
	title := getUserInput("Note Title:")

	content := getUserInput("Content:")

	return title, content
}