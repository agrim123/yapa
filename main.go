package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func GetPublicKey() {
	// Get User curretn Directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	// Read public key
	b, err := ioutil.ReadFile(usr.HomeDir + "/.ssh/id_rsa.pub")
	if err != nil {
		fmt.Print(err)
	}

	key := string(b)

	// Print public key
	fmt.Print(key)
}

func Ping() {
	result, _ := exec.Command("ping", "8.8.8.8", "-c 5", "-i 3", "-w 1").Output()

	if strings.Contains(string(result), "Destination Host Unreachable") {
		fmt.Println("I can't connect to internet!")
	} else {
		fmt.Println("We are online!")
	}

}

func main() {
	if len(os.Args) < 2 {
		panic("No args")
	}

	args := os.Args

	switch args[1] {
	case "key":
		GetPublicKey()
	case "ping":
		Ping()
	default:
		fmt.Println("Unkown Command")
	}
}
