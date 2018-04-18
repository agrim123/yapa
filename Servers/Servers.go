package Servers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var sshConfig *ssh.ClientConfig

type Server struct {
	Ip    string
	Users []string
}

type Config struct {
	Servers []Server
}

func GetPublicKey() {
	// Get User current Directory
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

func makeAuthFromAgent() (auth ssh.AuthMethod) {
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

	keyfile, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")

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
		auths = []ssh.AuthMethod{makeAuthFromAgent()}
	}

	sshConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func executeCmd(cmd, hostname string) string {
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

func ListServers() {
	var conf Config
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&conf); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	servers := conf.Servers
	var pointersToServers []*Server

	for i, x := range servers {
		pointersToServers = append(pointersToServers, &x)
		fmt.Println(i+1, ")", x.Ip, x.Users)
	}
}

func Poweroff() {
	cmd := exec.Command("sudo", "poweroff")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func Uptime(user, hostname string) {
	PrepareSSH(user)
	fmt.Println(executeCmd("uptime", hostname))
}
