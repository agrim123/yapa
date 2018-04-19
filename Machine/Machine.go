package Machine

import (
	"fmt"
	"log"
	"os/exec"
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
	RunCmd("sudo service apache2 stop")

}

func Count() {
	cmd := "ls | wc -l"

	fmt.Println("Running: " + cmd)

	RunCmd(cmd)
}
