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

func WriteToFile(path string, contents []byte) {
	ioutil.WriteFile(path, contents, 0644)
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

func ReadYapaConfig() (config *YapaConfig) {
	SetYapaConfigPath()

	b, err := ioutil.ReadFile(DefaultYapaConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(b, &config)

	return
}

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

func DisplayYapaConfig(config *YapaConfig) {
	fmt.Println(color.Blue("Username:"), config.Username)
	fmt.Println(color.Blue("System:"), config.System)
}

func SetupProfile() {
	oldConfig := ReadYapaConfig()

	if oldConfig != nil {
		fmt.Println("A yapa profile already exists:")

		DisplayYapaConfig(oldConfig)

		var k string

		fmt.Printf("Do you want to reset you profile? (Y/N): ")
		fmt.Scanf("%s", &k)

		if k == "Y" || k == "yes" || k == "y" {
			fmt.Println(color.Green("Overwriting your previous profile..."))
		} else {
			fmt.Println(color.Green("Going with the old profile."))
			return
		}
	} else {
		fmt.Println(color.Green("Setting up your new yapa profile..."))
	}

	var username string

	fmt.Printf(color.Blue("Username: "))
	fmt.Scanf("%s", &username)

	system, _ := os.Hostname()

	config := &YapaConfig{username, system}

	configJSON, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	WriteToFile(DefaultYapaConfigPath, configJSON)

	fmt.Println(color.Green("Saved Profile"))
}
