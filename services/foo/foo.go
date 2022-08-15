package main

import (
	"fmt"
	"net/http"
	"os"
	// "github.com/brandcodersteam/gotest-services/shared/"
)

var Thingy string

func init() {
	Thingy = "This is a thingy initializing..."
}

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	fmt.Println(Thingy)
	http.HandleFunc("/foo", HelloFooService)
	fmt.Println("Starting foo service...")
	http.ListenAndServe(":"+httpPort, nil)
}

func HelloFooService(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintf(w, "GET foo service endpoint: %s\n", r.URL.Path)
}
