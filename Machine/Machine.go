package Machine

import (
	"fmt"
	"log"
	"os/exec"
)

func Cool() {
	fmt.Println("Cooling off...")
	// Testing
	out, err := exec.Command("sudo", "service", "apache2", "stop").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
