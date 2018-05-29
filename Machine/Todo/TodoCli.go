package Todo

import (
	"../../Help"
	"fmt"
	"github.com/bclicn/color"
	"log"
)

func Cmd(args []string) {
	if len(args) == 0 {
		List()
		return
	}

	switch args[0] {
	case "list", "l":
		List()
	case "add", "a":
		Add()
	case "remove", "r":
		Remove(args[0:])
	case "complete", "c":
		Complete(args[0:])
	case "completed", "cp":
		ListCompletedTodos()
	case "incomplete", "incp":
		ListIncompleteTodos()
	case "h", "help":
		Help.BasicInfo()
		fmt.Print("COMMAND:")
		fmt.Println(Help.Todo())
	default:
		fmt.Println("COMMAND:")
		fmt.Println(Help.Todo())
		log.Fatal(color.Red("Unknown Command"))
	}
}
