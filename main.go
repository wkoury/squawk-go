package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", getSquawkEndpoint)
	http.ListenAndServe(":8088", nil)
}

func getSquawkEndpoint(w http.ResponseWriter, r *http.Request) {
	code := generate_squawk()
	io.WriteString(w, code)
}

func generate_squawk() string {
	return fmt.Sprintf("%d%d%d%d\n", rand.Intn(8), rand.Intn(8), rand.Intn(8), rand.Intn(8))
}
