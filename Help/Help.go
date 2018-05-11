package Help

import (
	"fmt"
	"github.com/bclicn/color"
	"log"
)

const VERSION = "v0.0.1"

func Global(endstring string) {
	const name = `NAME:
   yapa - Yet Another Personal Assistant
	`
	const usage = `USAGE:
   yapa [global options] command [command options] [arguments...]
	`
	version := fmt.Sprintf(`VERSION:
   %s
   `, VERSION)

	const commands = `COMMANDS:
   setup                        Setup a new yapa profile
   clean                        Clear all yapa settings
   count                        Count number of files/folders in directory
   key                          Get current user's public key
   ping [HOSTNAME]              Check if host is online
   list                         List all servers listed in config.json
   bye                          Shutdown system
   uptime [USER] [IP]           Display uptime of a server
   cool
   hackernews, hn               Display Hacker News
   scan                         Scan a hostname
   toss                         Flips a coin
   dice                         Roll a dice
   help, h                      Display help`
	const flags = `GLOBAL OPTIONS:
   -h, --help     Display help
	`

	fmt.Println(name)
	fmt.Println(usage)
	fmt.Println(version)
	fmt.Println(commands + Todo() + UserInfoHelp())
	fmt.Println(flags)

	if endstring != "" {
		log.Fatal(color.Red(endstring))
	}
}

func Todo() string {
	return `
   todo                         Show list of todo's
      list, l                   Show list of todo's
         completed, c           Show completed todo's
         incompleted, in        Show incomplete todo's
      remove, r [id]            Remove a todo from list
      add, a                    Add a new todo
      complete, c [id]          Mark a todo as completed
   `
}

func UserInfoHelp() string {
	return `
   all-users, allusr            List all users
   investigate, inv [username]  Get detail of the user specified
   `
}
