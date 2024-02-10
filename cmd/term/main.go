package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/supercaracal/scram-sha-256/pkg/pgpasswd"
	"golang.org/x/crypto/ssh/terminal"
)

func readViaTerminal(fd int) ([]byte, error) {
	fmt.Print("Raw password: ")
	passwd, err := terminal.ReadPassword(fd)
	fmt.Println()
	if err != nil {
		return nil, err
	}
	return passwd, nil
}

func readViaPipe() ([]byte, error) {
	r := bufio.NewReader(os.Stdin)
	passwd, err := r.ReadBytes('\n')
	if err == io.EOF {
		return passwd, nil
	} else if err != nil {
		return nil, err
	}
	return passwd[0 : len(passwd)-1], nil
}

func getRawPassword(args []string) ([]byte, error) {
	if len(args) > 1 {
		return []byte(args[1]), nil
	}

	fd := int(syscall.Stdin)
	if terminal.IsTerminal(fd) {
		return readViaTerminal(fd)
	}

	return readViaPipe()
}

func main() {
	rawPassword, err := getRawPassword(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(rawPassword) == 0 {
		fmt.Println("empty password")
		os.Exit(1)
	}

	encrypted, err := pgpasswd.Encrypt(rawPassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", encrypted)
	os.Exit(0)
}
