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
   help    Display help
   count   Count number of files/folders in directory
   key     Print current user public key
   ping    Check if online
   list    List all servers listed in config.json
   bye     Shutdown system
   uptime  Display uptime of a server
   cool
   scan    Scan a hostname   
	`

	fmt.Println(name)
	fmt.Println(usage)
	fmt.Println(version)
	fmt.Println(commands)

	if endstring != "" {
		log.Fatal(color.Red(endstring))
	}
}
