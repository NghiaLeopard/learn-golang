package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title," ","_") 
	fileName = strings.ToLower(fileName)  + ".json"

	json,err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName,json,0644)
}

func New(title, content string) (Note,error) {
	if title == "" || content == "" {
		return Note{},errors.New("title or content is empty")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	},nil
}

func (note Note) Print() {
	fmt.Printf("Note title: %v has following content:\n\n %v",note.Title,note.Content)
}