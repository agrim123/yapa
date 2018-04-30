package Todo

import (
	"fmt"
	"github.com/bclicn/color"
	"log"
	"strconv"
	"time"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Time        int64  `json:"time"`
	Completed   bool   `json:"completed"`
}

type Todos []*Todo

func List() {
	todos := ReadTodosFromFile()

	ViewTodosInList(todos)
}

func CreateNewTodo() *Todo {
	var title, description string

	fmt.Printf(color.Blue("Title: "))
	fmt.Scanf("%s", &title)

	fmt.Printf(color.Blue("Description: "))
	fmt.Scanf("%s", &description)

	todo := &Todo{title, description, time.Now().Unix(), false}

	return todo
}

func Add() {
	// Load all todos from file
	todos := ReadTodosFromFile()

	todo := CreateNewTodo()

	todos = append(todos, todo)

	// Save to file
	SaveTodosToFile(todos)

	fmt.Println(color.Green("Saved todo to list."))
}

func Remove(args []string) {
	if len(args) == 1 {
		log.Fatal(color.Red("Please provide id of todo to remove."))
	}

	// Load all todos from file
	todos := ReadTodosFromFile()

	CheckTodosLength(todos)

	todoToRemove, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatal(err)
	}

	if todoToRemove >= len(todos) {
		log.Fatal(color.Red("Couldn't find todo with that id."))
	}

	removedTodo := todos[todoToRemove]

	// Remove from file
	SaveTodosToFile(remove(todos, todoToRemove))

	fmt.Println(color.Green("Removed " + removedTodo.Title + " from store."))
}

func Complete(args []string) {
	if len(args) == 1 {
		log.Fatal(color.Red("Please provide id of todo to mark as completed."))
	}

	// Load all todos from file
	todos := ReadTodosFromFile()

	CheckTodosLength(todos)

	todoToComplete, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatal(err)
	}

	if todoToComplete >= len(todos) {
		log.Fatal(color.Red("Couldn't find todo with that id."))
	}

	todos[todoToComplete].Completed = true

	SaveTodosToFile(todos)

	fmt.Println(color.Green("Marked " + todos[todoToComplete].Title + " as completed."))
}

func ListCompletedTodos() {
	// Load all todos from file
	todos := ReadTodosFromFile()

	var filteredTodos []*Todo

	for _, v := range todos {
		if v.Completed == true {
			filteredTodos = append(filteredTodos, v)
		}
	}

	ViewTodosInList(filteredTodos)
}

func ListIncompleteTodos() {
	// Load all todos from file
	todos := ReadTodosFromFile()

	var filteredTodos []*Todo

	for _, v := range todos {
		if v.Completed == false {
			filteredTodos = append(filteredTodos, v)
		}
	}

	ViewTodosInList(filteredTodos)
}
