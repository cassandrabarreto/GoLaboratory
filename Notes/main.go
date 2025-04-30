package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	err = userNote.Save()

	if err != nil{
		fmt.Println("Error Saving.")
		return 
		
	}

	fmt.Println("Note Saved.")
	
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

/* Function that returns two parts: title and Note.*/
func getNoteData()(string, string){
	title := getUserInput("Note Title:")

	content := getUserInput("Content:")

	return title, content
}