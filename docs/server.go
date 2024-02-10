package main

import (
	"fmt"
	"net/http"
)

const (
	appName = "scram-sha-256"
)

var (
	staticFiles = []string{
		"encrypt.wasm",
		"favicon.ico",
		"script.js",
		"style.css",
		"wasm_exec.js",
	}
)

func makeStaticHandler(f string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, f)
	}
}

func main() {
	http.HandleFunc("/", makeStaticHandler("index.html"))

	for _, f := range staticFiles {
		http.HandleFunc(fmt.Sprintf("/%s/%s", appName, f), makeStaticHandler(f))
	}

	http.ListenAndServe(":3000", nil)
}
