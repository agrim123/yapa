package Todo

import (
	"../../Utility"
	"encoding/json"
	"fmt"
	"github.com/bclicn/color"
	"log"
	"strconv"
)

func ReadTodosFromFile() Todos {
	file := Utility.ReadFile(Utility.DefaultYapaTodoJSONPath)

	var todos Todos
	json.Unmarshal(file, &todos)

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

	Utility.WriteToFile(Utility.DefaultYapaTodoJSONPath, todosJSON)
}

func ViewTodosInList(todos Todos) {
	CheckTodosLength(todos)

	for k, v := range todos {
		fmt.Println(strconv.Itoa(k)+")", v.Title, v.Description, v.Time, v.Completed)
	}
}

func RemoveTodo(s Todos, i int) Todos {
	return append(s[:i], s[i+1:]...)
}
