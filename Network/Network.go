package Network

import (
	"fmt"
	"github.com/bclicn/color"
	"log"
	"net"
	"os/exec"
	"strings"
)

func Ping(args []string) {
	hostname := "8.8.8.8"
	message := "We are online!"
	err := "I can't connect to internet!"

	if len(args) > 1 {
		hostname = args[1]
		message = hostname + " is up."
		err = hostname + " is down."
	}

	result, _ := exec.Command("ping", hostname, "-c 5", "-i 3", "-w 1").Output()

	if strings.Contains(string(result), "Destination Host Unreachable") {
		log.Fatal(color.Red(err))
	} else {
		fmt.Println(color.Green(message))
	}
}

func Scan(args []string) {
	if len(args) < 2 {
		log.Fatal("Please supply a hostname to scan.")
	}

	hostname := net.ParseIP(args[1])

	if hostname == nil {
		log.Fatal(color.Red("Please provide a correct IP address."))
	}

	// Todo: Scan a host
}

func SpeedTest() {
	// Ping()

	fmt.Println("Testing Speed...")

}
