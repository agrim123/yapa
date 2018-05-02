package Servers

import (
	"../Utility"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bclicn/color"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
)

var sshConfig *ssh.ClientConfig

type Server struct {
	Ip    string
	Users []string
}

type Servers []*Server

func GetPublicKey() {
	// Read public key
	b, err := ioutil.ReadFile(Utility.UserPublicKeyPath())
	if err != nil {
		fmt.Print(err)
	}

	key := string(b)

	// Print public key
	fmt.Println(key)
}

func MakeAuthFromAgent() (auth ssh.AuthMethod) {
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Fatal(err)
	}
	ag := agent.NewClient(conn)
	auth = ssh.PublicKeysCallback(ag.Signers)
	return
}

func PrepareSSH(user string) {
	var auths []ssh.AuthMethod
	var privateKey ssh.Signer

	keyfile, err := ioutil.ReadFile(Utility.UserPrivateKeyPath())

	if err == nil {
		key, err := ssh.ParsePrivateKey(keyfile)
		if err == nil {
			privateKey = key
		}
	}

	if privateKey != nil {
		auths = []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		}
	} else {
		auths = []ssh.AuthMethod{MakeAuthFromAgent()}
	}

	sshConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func executeCmd(cmd, hostname string) string {
	fmt.Println(color.Green("Running " + cmd + " on " + sshConfig.User + "@" + hostname))

	client, err := ssh.Dial("tcp", hostname+":22", sshConfig)
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln("Failed to create session:", err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = os.Stderr

	if err := session.Run(cmd); err != nil {
		log.Fatalln("Failed to run:", err)
	}

	return stdoutBuf.String()
}

func GetServers() (servers []Server) {
	configFile, err := os.Open(Utility.DefaultYapaServerConfigPath)
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&servers); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	return
}

func ListServers() {
	servers := GetServers()

	for i, x := range servers {
		fmt.Println(i, ")", x.Ip, x.Users)
	}
}

func Poweroff() {
	cmd := exec.Command("sudo", "poweroff")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func ParseHostname(hostname string) string {
	parsed := net.ParseIP(hostname)

	if parsed == nil {
		// Look in store
		servers := GetServers()
		index, _ := strconv.Atoi(hostname)
		return servers[index].Ip
	}

	return parsed.String()
}

func Uptime(args []string) {
	if len(args) < 3 {
		log.Fatal(color.Red("Please supply a user and hostname to view uptime of server."))
	}

	user := args[1]
	hostname := args[2]

	PrepareSSH(user)
	fmt.Println(executeCmd("uptime", ParseHostname(hostname)))
}
