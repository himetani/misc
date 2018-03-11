package main

import (
	"fmt"
	"io"
	"log"
	"sync"

	"golang.org/x/crypto/ssh"
)

func MuxShell(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()
	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				close(in)
				close(out)
				return
			}
			t += n
			if buf[t-2] == '$' { //assuming the $PS1 == 'sh-4.3$ '
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}

func main() {
	config := &ssh.ClientConfig{
		User:            "test",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password("fakepass"),
		},
	}
	client, err := ssh.Dial("tcp", "127.0.0.1:2222", config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer client.Close()
	session, err := client.NewSession()

	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatal(err)
	}

	w, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	r, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}
	in, out := MuxShell(w, r)
	if err := session.Start("/bin/sh"); err != nil {
		log.Fatal(err)
	}
	<-out //ignore the shell output
	in <- "echo 'fakepass' | sudo -S su - vagrant"
	fmt.Printf("ls output: %s\n", <-out)

	in <- "whoami"
	fmt.Printf("whoami: %s\n", <-out)

	in <- "ls -ltr"
	fmt.Printf("whoami: %s\n", <-out)

	in <- "pwd"
	fmt.Printf("pwd: %s\n", <-out)

	in <- "exit"
	session.Wait()
}
