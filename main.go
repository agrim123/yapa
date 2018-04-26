package main

import (
	"./Machine"
	"./Machine/Todo"
	"./Network"
	"./Servers"
	"./Utility"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Forever() {
	exitCmd := []string{"exit", "quit", "e", "q"}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// strip '\n'
		text = strings.Replace(text, "\n", "", -1)

		cmds := strings.Split(text, " ")

		if Utility.ArrayContains(cmds[0], exitCmd) == true {
			fmt.Println("Exiting forever mode.")
			break
		}

		ParseArgs(cmds)
	}
}

func ParseArgs(args []string) {
	switch args[0] {
	case "help", "h", "-h", "--help":
		Utility.Help("")
	case "setup":
		Machine.Setup()
	case "clean":
		Machine.Clean()
	case "key":
		Servers.GetPublicKey()
	case "ping":
		Network.Ping(args[0:])
	case "list":
		Servers.ListServers()
	case "bye":
		Servers.Poweroff()
	case "uptime":
		Servers.Uptime(args[0:])
	case "cool":
		Machine.Cool()
	case "scan":
		Network.Scan(args[0:])
	case "count":
		Machine.Count()
	case "speedtest":
		Network.SpeedTest()
	case "hackernew", "hn":
		Network.HackerNews()
	case "toss":
		Machine.Toss()
	case "dice":
		Machine.Dice()
	case "todo":
		Todo.Cmd(args[0:])
	case "forever":
		Forever()
	default:
		Utility.Help("Unkown Command")
	}
}

func main() {
	if len(os.Args) < 2 {
		Utility.Help("No arguments specified.")
	}

	args := os.Args

	ParseArgs(args[1:])
}
