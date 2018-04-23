package Machine

import (
	"fmt"
	"github.com/bclicn/color"
	"log"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	title       string `json:"title"`
	description string `json:"description"`
	time        int64  `json:"time"`
	completed   bool   `json:"completed"`
}

var Todos []*Todo

func ReadFromFile() {

}

func SaveToFile() {

}

func List() {
	if len(Todos) == 0 {
		fmt.Println(color.Red("No todos found"))
		return
	}

	for k, v := range Todos {
		fmt.Println(k, v)
	}
}

func Add() {
	var title, description string

	fmt.Printf("Title: ")
	fmt.Scanf("%s", &title)

	fmt.Printf("Description: ")
	fmt.Scanf("%s", &description)

	todo := new(Todo)

	todo.title = title
	todo.description = description
	todo.time = time.Now().Unix()
	todo.completed = false

	Todos = append(Todos, todo)

	// Todo: Save to file

	fmt.Println(color.Blue("Saved todo to list."))
}

func Remove() {
	if len(os.Args) < 4 {
		log.Fatal(color.Red("Please provide id of todo to remove."))
	}

	todoToRemove, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	if todoToRemove > len(Todos) {
		log.Fatal(color.Red("Couldn't find todo with that id."))
	}

	// Remove from file

}

func Complete() {
	if len(os.Args) < 4 {
		log.Fatal(color.Red("Please provide id of todo to mark as completed."))
	}

	todoToComplete, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	if todoToComplete > len(Todos) {
		log.Fatal(color.Red("Couldn't find todo with that id."))
	}

}

func TodoCmd() {
	args := os.Args

	if len(args) < 3 {
		List()
		return
	}

	switch args[2] {
	case "list", "l":
		List()
	case "add", "a":
		Add()
	case "remove", "r":
		Remove()
	case "complete", "c":
		Complete()
	default:
		log.Fatal(color.Red("Unknown Command"))
	}
}
