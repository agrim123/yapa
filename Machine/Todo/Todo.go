package Todo

import (
	"../../Utility"
	"encoding/json"
	"fmt"
	"github.com/bclicn/color"
	"io/ioutil"
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

func ReadTodosFromFile() Todos {
	Utility.CreateFile(Utility.DefaultYapaTodoJSONPath, "Found "+color.Blue("todo.json"), "Todo store does not exist. Creating a new one...")

	b, err := ioutil.ReadFile(Utility.DefaultYapaTodoJSONPath)
	if err != nil {
		log.Fatal(err)
	}

	var todos Todos
	json.Unmarshal(b, &todos)

	return todos
}

func CheckTodosLength(todos Todos) {
	if len(todos) == 0 {
		log.Fatal(color.Red("No todos found"))
	}
}

func SaveTodosToFile(todos Todos) {
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(Utility.DefaultYapaTodoJSONPath, todosJSON, 0644)
}

func ViewTodosInList(todos Todos) {
	CheckTodosLength(todos)

	for k, v := range todos {
		fmt.Println(strconv.Itoa(k)+")", v.Title, v.Description, v.Time, v.Completed)
	}
}

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

	todo := new(Todo)

	todo.Title = title
	todo.Description = description
	todo.Time = time.Now().Unix()
	todo.Completed = false

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

func remove(s Todos, i int) Todos {
	return append(s[:i], s[i+1:]...)
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

func Cmd(args []string) {
	if len(args) == 1 {
		List()
		return
	}

	switch args[1] {
	case "list", "l":
		List()
	case "add", "a":
		Add()
	case "remove", "r":
		Remove(args[1:])
	case "complete", "c":
		Complete(args[1:])
	case "completed", "cp":
		ListCompletedTodos()
	case "incomplete", "incp":
		ListIncompleteTodos()
	default:
		fmt.Print("COMMAND:")
		fmt.Println(Utility.TodoHelp())
		log.Fatal(color.Red("Unknown Command"))
	}
}
