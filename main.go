package main

import (
	"./Machine"
	"./Network"
	"./Servers"
	"./Utility"
	"os"
)

func ParseArgs(command string) {
	switch command {
	case "help":
		Utility.Help("")
	case "key":
		Servers.GetPublicKey()
	case "ping":
		Network.Ping()
	case "list":
		Servers.ListServers()
	case "bye":
		Servers.Poweroff()
	case "uptime":
		Servers.Uptime()
	case "cool":
		Machine.Cool()
	case "scan":
		Network.Scan()
	case "count":
		Machine.Count()
	default:
		Utility.Help("Unkown Command")
	}
}

func main() {
	if len(os.Args) < 2 {
		Utility.Help("No arguments specified.")
	}

	args := os.Args

	ParseArgs(args[1])
}
