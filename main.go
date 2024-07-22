package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
)

func main() {
	title,content := getNote()

	userNote,err := note.New(title,content)

	if err != nil {
		fmt.Println(err)
	}

	userNote.Print()

	err1 := userNote.Save()

	if err1 != nil {
		return
	}

}

func getNote() (string ,string) {
	title := setInputUser("Please enter title: ")
	content := setInputUser("Please enter content: ")

	return title,content
}

func setInputUser(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text,err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")

	return text
}