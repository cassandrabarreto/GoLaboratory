package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Create constructor
func NewNote(title, content string) (Note, error){

	if title == "" || content == ""{
		//Return dummy Note struct.
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title : title,
		Content : content,
		CreatedAt : time.Now(),
	}, nil
}

// Create 'class' (struct ) method
func (note Note) Display() {
	fmt.Printf("Your note is titled: %v\n has the following content: %v\n", note.Title, note.Content)
}

func (note Note)Save() error{
	filename := strings.ReplaceAll(note.Title," ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note)

	if err != nil{
		return err
	}
	return os.WriteFile(filename, json, 0644)

}