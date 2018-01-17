package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func PublicKey() {
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

	str := string(b)

	// Print public key
	fmt.Print(str)
}

func main() {
	if len(os.Args) < 2 {
		panic("No args")
	}

	args := os.Args

	switch args[1] {
	case "key":
		PublicKey()
	default:
		fmt.Println("Unkown Command")
	}
}
