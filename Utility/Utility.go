package Utility

import (
	"encoding/json"
	"fmt"
	"github.com/bclicn/color"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

const VERSION = "v0.0.1"

func Help(endstring string) {
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
   setup                     Setup a new yapa profile
   clean                     Clear all yapa settings
   count                     Count number of files/folders in directory
   key                       Get current user's public key
   ping [HOSTNAME]           Check if host is online
   list                      List all servers listed in config.json
   bye                       Shutdown system
   uptime [USER] [IP]        Display uptime of a server
   cool
   hackernews, hn            Display Hacker News
   scan                      Scan a hostname
   toss                      Flips a coin
   dice                      Roll a dice
   help, h                   Display help`
	const flags = `GLOBAL OPTIONS:
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

func CreateFileIfNotExists(path string, found string, notfound string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		fmt.Println(notfound)

		var file, err = os.Create(path)
		fmt.Println("Created " + color.Blue(path))

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
	} else {
		fmt.Println(found)
	}
}

func CreateDirIfNotExists(path string, found string, notfound string, perm int) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(notfound)
		os.Mkdir(path, os.FileMode(perm))
		fmt.Println("Created " + color.Blue(path))
	} else {
		fmt.Println(found)
	}
}

func UserHomeDir() string {
	// Get User current Directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	return usr.HomeDir
}

func UserPublicKeyPath() string {
	return UserHomeDir() + "/.ssh/id_rsa.pub"
}

func UserPrivateKeyPath() string {
	return UserHomeDir() + "/.ssh/id_rsa"
}

type YapaConfig struct {
	Username string `json:"username"`
	System   string `json:"system"`
}

func ReadYapaConfig() {
	SetYapaConfigPath()

	b, err := ioutil.ReadFile(DefaultYapaConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	var config YapaConfig
	json.Unmarshal(b, &config)

	fmt.Println(config)
}

const (
	YapaDir = ".yapa"

	ConfigJSON = "config.json"

	TodoJSON = "todo.json"

	ServersJSON = "servers.json"
)

// Variables related to yapa
var (
	DefaultYapaDir = UserHomeDir() + "/" + YapaDir

	DefaultYapaConfigPath = DefaultYapaDir + "/" + ConfigJSON

	DefaultYapaTodoJSONPath = DefaultYapaDir + "/" + TodoJSON

	DefaultYapaServerConfigPath = DefaultYapaDir + "/" + ServersJSON
)

func SetYapaDir() {
	CreateDirIfNotExists(DefaultYapaDir, "Found "+color.Blue(YapaDir), "Default yapa directory doesnot exist. Creating a new one...", 0775)
}

func SetYapaConfigPath() {
	CreateFileIfNotExists(DefaultYapaConfigPath, "Found "+color.Blue(ConfigJSON), "Default config doesnot exist. Creating a new one...")
}

func SetYapaServerConfigPath() {
	CreateFileIfNotExists(DefaultYapaServerConfigPath, "Found "+color.Blue(ServersJSON), "Default server config doesnot exist. Creating a new one...")
}

func SetYapaTodoJSONPath() {
	CreateFileIfNotExists(DefaultYapaTodoJSONPath, "Found "+color.Blue(TodoJSON), "Todo store does not exist. Creating a new one...")
}
