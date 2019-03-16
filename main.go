package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	//cli
	//connect to server
	log.Println("Starting ssh service")
	host := ""
	user := ""
	// var hostkey ssh.PublicKey

	// Get private key
	key, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	if err != nil {
		log.Printf("unable to read private key: %v", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Printf("unable to parse private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatal("unable to connect: %v", err)
	}
	defer client.Close()

	//get ls
	session, _ := client.NewSession()
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("pwd")
	log.Println(stdoutBuf.String())

	//create file
	//test
}
