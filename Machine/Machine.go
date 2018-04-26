package Machine

import (
	"../Utility"
	"fmt"
	"github.com/bclicn/color"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func RunCmd(cmd string) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func Cool() {
	fmt.Println("Cooling off...")
	// Todo
	RunCmd("sudo service apache2 stop")
}

func Count() {
	cmd := "ls | wc -l"

	fmt.Println(color.Blue("Running: ") + cmd)

	RunCmd(cmd)
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
	Utility.CreateDir(Utility.DefaultYapaDir, "Found "+color.Blue(".yapa"), "Default directory doesnot exist. Creating a new one...", 0755)

	Utility.CreateFile(Utility.DefaultYapaConfigPath, "Found "+color.Blue(".yapa/config.json"), "Default config doesnot exist. Creating a new one...")

	SetupConfig()

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(b))
}

func SetupConfig() {

}

func Clean() {
	var err = os.RemoveAll(Utility.UserHomeDir() + "/.yapa")

	if err != nil {
		log.Fatal(err)
	}
}
