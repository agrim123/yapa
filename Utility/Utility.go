package Utility

import (
	"fmt"
	"github.com/bclicn/color"
	"log"
)

const VERSION = "v0.0.1"

func Help(endstring string) {
	const name = `NAME:
   yapa - Yet Another Personal Assistant
	`
	const usage = `USAGE:
   yapa [command]
	`
	version := fmt.Sprintf(`VERSION:
   %s
   `, VERSION)

	const commands = `COMMANDS:
   help, h                   Display help
   count                     Count number of files/folders in directory
   key                       Print current user public key
   ping                      Check if online
   list                      List all servers listed in config.json
   bye                       Shutdown system
   uptime                    Display uptime of a server
   cool
   hackernews, hn            Display Hacker News
   scan                      Scan a hostname
   toss                      Flips a coin
   dice                      Roll a dice`
	const flags = `FLAGS:
   -h, --help     Display help
	`

	fmt.Println(name)
	fmt.Println(usage)
	fmt.Println(version)
	fmt.Println(commands + TodoHelp())
	fmt.Println(flags)

	if endstring != "" {
		log.Fatal(color.Red(endstring))
	}
}

func TodoHelp() string {
	return `
   todo                      Show list of todo's
      list, l                Show list of todo's
         completed, c        Show completed todo's  
         incompleted, in     Show incomplete todo's
      remove, r [id]         Remove a todo from list
      add, a                 Add a new todo
      complete, c [id]       Mark a todo as completed
	`
}

func ArrayContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
