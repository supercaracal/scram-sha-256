package main

// @see https://github.com/postgres/postgres/blob/c30f54ad732ca5c8762bb68bbe0f51de9137dd72/src/interfaces/libpq/fe-auth.c#L1167-L1285
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/interfaces/libpq/fe-auth-scram.c#L868-L905
// @see https://github.com/postgres/postgres/blob/c30f54ad732ca5c8762bb68bbe0f51de9137dd72/src/port/pg_strong_random.c#L66-L96
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/common/scram-common.c#L160-L274
// @see https://github.com/postgres/postgres/blob/e6bdfd9700ebfc7df811c97c2fc46d7e94e329a2/src/common/scram-common.c#L27-L85

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"syscall"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	saltSize     = 16
	digestLen    = 32
	iterationCnt = 4096
)

func genSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func readRawPassword(fd int) ([]byte, error) {
	input, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func encodeB64(src []byte) (dst []byte) {
	dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return
}

func getHashSum(key, msg []byte) []byte {
	h := hmac.New(sha256.New, key)
	_, _ = h.Write(msg)
	return h.Sum(nil)
}

func getStoredKey(clientKey []byte) (storedKey []byte) {
	key := sha256.Sum256(clientKey)
	storedKey = make([]byte, 0, len(key))
	for _, b := range key {
		storedKey = append(storedKey, b)
	}
	return
}

func encryptPassword(rawPassword, salt []byte, iter, keyLen int) string {
	digestKey := pbkdf2.Key(rawPassword, salt, iter, keyLen, sha256.New)
	clientKey := getHashSum(digestKey, []byte("Client Key"))
	storedKey := getStoredKey(clientKey)
	serverKey := getHashSum(digestKey, []byte("Server Key"))

	return fmt.Sprintf("SCRAM-SHA-256$%d:%s$%s:%s",
		iter,
		string(encodeB64(salt)),
		string(encodeB64(storedKey)),
		string(encodeB64(serverKey)),
	)
}

func main() {
	fmt.Print("Raw password: ")
	rawPassword, err := readRawPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(rawPassword) == 0 {
		fmt.Println("empty")
		os.Exit(1)
	}

	salt, err := genSalt(saltSize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", encryptPassword(rawPassword, salt, iterationCnt, digestLen))
	os.Exit(0)
}
