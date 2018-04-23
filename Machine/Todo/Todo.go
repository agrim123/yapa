package Todo

import (
	"encoding/json"
	"fmt"
	"github.com/bclicn/color"
	"io/ioutil"
	"log"
	"os"
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

func GetTodoFilePath() string {
	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir + "/Machine/todo.json"
}

func ReadTodosFromFile() Todos {
	// todo: Handle case when todo.json doesnot exist
	b, err := ioutil.ReadFile(GetTodoFilePath())
	if err != nil {
		log.Fatal(err)
	}

	var todos Todos
	json.Unmarshal(b, &todos)

	return todos
}

func CheckTodosLength(todos Todos) {
	if len(todos) == 0 {
		fmt.Println(color.Red("No todos found"))
		os.Exit(1)
	}
}

func SaveTodosToFile(todos Todos) {
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(GetTodoFilePath(), todosJSON, 0644)
}

func List() {
	todos := ReadTodosFromFile()

	CheckTodosLength(todos)

	for k, v := range todos {
		fmt.Println(strconv.Itoa(k)+")", v.Title, v.Description, v.Time, v.Completed)
	}
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

func Remove() {
	if len(os.Args) < 4 {
		log.Fatal(color.Red("Please provide id of todo to remove."))
	}

	// Load all todos from file
	todos := ReadTodosFromFile()

	CheckTodosLength(todos)

	todoToRemove, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	if todoToRemove > len(todos) {
		log.Fatal(color.Red("Couldn't find todo with that id."))
	}

	removedTodo := todos[todoToRemove]

	// Remove from file
	SaveTodosToFile(remove(todos, todoToRemove))

	fmt.Println(color.Green("Removed " + removedTodo.Title + " from store."))
}

func Complete() {
	if len(os.Args) < 4 {
		log.Fatal(color.Red("Please provide id of todo to mark as completed."))
	}

	// Load all todos from file
	todos := ReadTodosFromFile()

	CheckTodosLength(todos)

	todoToComplete, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	if todoToComplete > len(todos) {
		log.Fatal(color.Red("Couldn't find todo with that id."))
	}

	todos[todoToComplete].Completed = true

	SaveTodosToFile(todos)

	fmt.Println(color.Green("Marked " + todos[todoToComplete].Title + " as completed."))
}

func ListCompletedTodod() {

}

func ListIncompleteTodos() {

}

func Cmd() {
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
