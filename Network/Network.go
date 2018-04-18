package Network

import (
	"fmt"
	"github.com/bclicn/color"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func Ping() {
	result, _ := exec.Command("ping", "8.8.8.8", "-c 5", "-i 3", "-w 1").Output()

	if strings.Contains(string(result), "Destination Host Unreachable") {
		fmt.Println(color.Red("I can't connect to internet!"))
	} else {
		fmt.Println(color.Green("We are online!"))
	}
}

func Scan() {
	if len(os.Args) < 3 {
		log.Fatal("Please supply a hostname to scan.")
	}

	hostname := net.ParseIP(os.Args[2])

	if hostname == nil {
		log.Fatal(color.Red("Please provide a correct IP address."))
	}

	// Scan a host
}
