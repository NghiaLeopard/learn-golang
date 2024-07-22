package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
	"example.com/main/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Print()
}

func main() {
	title,content := getNote()
	text := setInputUser("Please enter text: ")
	printSomething(1)
	printSomething(1.5)
	printSomething("Beautifully ✨")


	userNote,err3:= note.New(title,content)
	userTodo,err := todo.New(text)

	if err != nil || err3 != nil {
		fmt.Println(err)
	}
	err1 := outputTable(userNote)

	err2 := outputTable(userTodo)
	
	if err1 != nil || err2 != nil {
		return 
	}
}

func printSomething(value interface{}) {
	inputVal,ok := value.(int)

	if ok {
	fmt.Printf("Integer: %d\n",inputVal)

	}
	// switch value.(type) {
	// case int: 
	// fmt.Printf("Integer: %d\n",value)
	// case float64: 
	// fmt.Printf("Integer: %f\n",value)
	// case string:
	// fmt.Printf("string: %s\n",value)
	// }
}

func outputTable(data outputtable) error {
	data.Print()
	err := saveData(data)

	return err
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil  {
		return errors.New("Data is saved fail")
	}

	fmt.Println("Data is saved successfully ✨")

	return nil
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