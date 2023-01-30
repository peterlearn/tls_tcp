package main

import (
	"fmt"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS("127.0.0.1:8003", "/var/run/test/server.crt", "/var/run/test/server.key", nil)
}
