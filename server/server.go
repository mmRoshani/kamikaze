package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n%s\n", s.User(), s.Command()))
	})

	log.Println("starting ssh server on port 2030")
	log.Fatal(ssh.ListenAndServe(":2030", nil))
}
