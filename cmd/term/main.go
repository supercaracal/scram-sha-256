package main

// @see https://github.com/postgres/postgres/blob/c30f54ad732ca5c8762bb68bbe0f51de9137dd72/src/interfaces/libpq/fe-auth.c#L1167-L1285
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/interfaces/libpq/fe-auth-scram.c#L868-L905
// @see https://github.com/postgres/postgres/blob/c30f54ad732ca5c8762bb68bbe0f51de9137dd72/src/port/pg_strong_random.c#L66-L96
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/common/scram-common.c#L160-L274
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/common/scram-common.c#L27-L85

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
