package Network

import (
	"fmt"
	"os/exec"
	"strings"
)

func Ping() {
	result, _ := exec.Command("ping", "8.8.8.8", "-c 5", "-i 3", "-w 1").Output()

	if strings.Contains(string(result), "Destination Host Unreachable") {
		fmt.Println("I can't connect to internet!")
	} else {
		fmt.Println("We are online!")
	}
}

func Scan(hostname string) {
	fmt.Println(hostname)
}
