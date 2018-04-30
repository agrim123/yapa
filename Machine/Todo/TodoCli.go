package Todo

import (
	"../../Utility"
	"fmt"
	"github.com/bclicn/color"
	"log"
)

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
