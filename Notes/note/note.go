package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct{
	title string
	content string
	createdAt time.Time
}

// Create constructor
func NewNote(title, content string) (Note, error){

	if title == "" || content == ""{
		//Return dummy Note struct.
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title : title,
		content : content,
		createdAt : time.Now(),
	}, nil
}

// Create 'class' (struct ) method
func (note Note) Display() {
	fmt.Printf("Your note titled: %v has the following content: %v", note.title, note.content)
}