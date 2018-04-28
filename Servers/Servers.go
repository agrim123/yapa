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
	// Read public key
	b, err := ioutil.ReadFile(Utility.UserPublicKeyPath())
	if err != nil {
		fmt.Print(err)
	}

	key := string(b)

	// Print public key
	fmt.Println(key)
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
	configFile, err := os.Open(Utility.DefaultYapaServerConfigPath)
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

func Uptime(args []string) {
	if len(args) < 3 {
		log.Fatal(color.Red("Please supply a user and hostname to view uptime of server."))
	}

	user := args[1]
	hostname := args[2]

	PrepareSSH(user)
	fmt.Println(executeCmd("uptime", hostname))
}
