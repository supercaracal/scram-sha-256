package main

// @see https://github.com/postgres/postgres/blob/c30f54ad732ca5c8762bb68bbe0f51de9137dd72/src/interfaces/libpq/fe-auth.c#L1167-L1285
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/interfaces/libpq/fe-auth-scram.c#L868-L905
// @see https://github.com/postgres/postgres/blob/c30f54ad732ca5c8762bb68bbe0f51de9137dd72/src/port/pg_strong_random.c#L66-L96
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/common/scram-common.c#L160-L274
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/common/scram-common.c#L27-L85

import (
	"fmt"
	"os"
	"syscall"

	"github.com/supercaracal/scram-sha-256/pkg/pgpasswd"
	"golang.org/x/crypto/ssh/terminal"
)

func readRawPassword(fd int) ([]byte, error) {
	input, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func main() {
	var rawPassword []byte

	if len(os.Args) > 1 {
		rawPassword = []byte(os.Args[1])
	} else {
		fmt.Print("Raw password: ")
		passwd, err := readRawPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rawPassword = passwd
		fmt.Println()
	}

	if len(rawPassword) == 0 {
		fmt.Println("empty password")
		os.Exit(1)
	}

	if password, err := pgpasswd.Encrypt(rawPassword); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", password)
		os.Exit(0)
	}
}
