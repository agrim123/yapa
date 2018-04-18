package main

import (
	"./Servers"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("No args")
	}

	args := os.Args

	switch args[1] {
	case "key":
		Servers.GetPublicKey()
	case "ping":
		Servers.Ping()
	case "list":
		Servers.ListServers()
	case "bye":
		Servers.Poweroff()
	case "uptime":
		Servers.Uptime("", "")
	default:
		fmt.Println("Unkown Command")
	}
}
