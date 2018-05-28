package main

import (
	"./Help"
	"./Machine"
	"./Machine/Todo"
	"./Network"
	"./Servers"
	"./Utility"
	"bufio"
	"fmt"
	"github.com/bclicn/color"
	"os"
	"strings"
)

func Forever() {
	exitCmd := []string{"exit", "quit", "e", "q"}
	fmt.Println(color.Blue("You have entered forever mode. To quit at anytime, type any of the following commands: exit, quit, e, q"))

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// strip '\n'
		text = strings.Replace(text, "\n", "", -1)

		cmds := strings.Split(text, " ")

		if len(cmds) < 1 || len(cmds[0]) == 0 {
			continue
		}

		if Utility.ArrayContains(cmds[0], exitCmd) == true {
			fmt.Println("Exiting forever mode.")
			break
		}

		ParseArgs(cmds, true)
	}
}

func ParseArgs(args []string, forever bool) {
	switch args[0] {
	case "help", "h", "-h", "--help":
		Help.Global("", forever)
	case "setup", "init":
		Machine.Setup()
	case "clean":
		Machine.Clean()
	case "show", "profile":
		Machine.Profile()
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
	case "server":
		Servers.Init(args[0:])
	case "cool":
		Machine.Cool()
	case "scan":
		Network.Scan(args[0:])
	case "count":
		Machine.Count()
	case "speedtest":
		Network.SpeedTest()
	case "hackernews", "hn":
		Network.HackerNews()
	case "toss":
		Machine.Toss()
	case "dice":
		Machine.Dice()
	case "todo":
		Todo.Cmd(args[0:])
	case "forever":
		Forever()
	case "all-users", "allusr":
		Machine.AllUsers()
	case "investigate", "inv":
		Machine.InvestigateUser(args[1])
	case "ip":
		Network.PublicIP()
	default:
		Help.Global("Unkown Command", forever)
	}
}

func main() {
	if len(os.Args) < 2 {
		Help.Global("No arguments specified.", false)
	}

	args := os.Args

	ParseArgs(args[1:], false)
}
