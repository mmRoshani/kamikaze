package main

import (
	"os"

	"os/user"

	"golang.org/x/crypto/ssh"
)

func main() {

	//! fetch current username
	currentUser, err := user.Current()
	if err != nil {
		panic(err.Error())
	}

	//! windows
	key, err := os.ReadFile(currentUser.HomeDir + "\\.ssh\\id_rsa")
	//! mac
	// key, err := os.ReadFile("/home/user/.ssh/id_rsa")
	//! linux
	if err != nil {
		panic("ERROR> unable to read private key|\n" + err.Error())
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		panic("ERROR> Failed to dial|\n" + err.Error())
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic("ERROR> Failed to create session|\n" + err.Error())
	}
	defer session.Close()

	//! run shell

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		panic("ERROR> Failed to run|\n" + err.Error())
	}

	if err := session.Wait(); err != nil {
		panic(err)
	}
}
