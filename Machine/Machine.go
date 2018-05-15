package Machine

import (
	"../Utility"
	"fmt"
	"github.com/bclicn/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Cool() {
	fmt.Println("Cooling off...")
	// Todo
	Utility.RunCmd("sudo service apache2 stop")
}

func Count() {
	Utility.RunCmd("ls -la | wc -l")
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func Toss() {
	rand.Seed(time.Now().UnixNano())
	num := random(1, 512)

	if num%2 == 1 {
		fmt.Println(color.Blue("Tails"))
	} else {
		fmt.Println(color.Blue("Heads"))
	}
}

func Dice() {
	rand.Seed(time.Now().UnixNano())
	num := random(1, 6)

	fmt.Println(color.Blue(strconv.Itoa(num)))
}

func Setup() {
	Utility.SetYapaDir()

	Utility.SetYapaConfigPath()

	Utility.SetYapaServerConfigPath()

	Utility.SetYapaTodoJSONPath()

	Utility.SetupProfile()
}

func Clean() {
	var err = os.RemoveAll(Utility.UserHomeDir() + "/.yapa")
	fmt.Println(color.Blue("Removing .yapa..."))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(color.Green("Cleaned up all config."))
}

func Profile() {
	config := Utility.ReadYapaConfig()

	Utility.DisplayYapaConfig(config)
}

func AllUsers() {
	Utility.RunCmd("awk -F':' '{ print $1}' /etc/passwd")
}

func InvestigateUser(user string) {
	if user == "" {
		log.Fatal(color.Red("No user to investigate."))
	}

	Utility.RunCmd("id " + user)

	Utility.RunCmd("groups " + user)

	Utility.RunCmd("< /etc/passwd grep " + user)

	Utility.RunCmd("< /etc/group grep " + user)

	Utility.RunCmd("finger " + user)
}
