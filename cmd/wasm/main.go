package main

import (
	"syscall/js"

	"github.com/supercaracal/scram-sha-256/pkg/pgpasswd"
)

func handleButton(this js.Value, inputs []js.Value) interface{} {
	document := js.Global().Get("document")
	input := document.Call("getElementById", "raw-password")
	result := document.Call("getElementById", "encrypted-password")

	rawPassword := input.Get("value").String()

	if encrypted, err := pgpasswd.Encrypt([]byte(rawPassword)); err != nil {
		result.Set("value", err.Error())
	} else {
		result.Set("value", encrypted)
	}

	return nil
}

func main() {
	button := js.Global().Get("document").Call("getElementById", "encryption-button")
	button.Call("addEventListener", "click", js.FuncOf(handleButton))
	select {}
}
