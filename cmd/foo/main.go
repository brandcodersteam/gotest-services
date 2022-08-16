package main

import (
	"fmt"
	"gotest-services/pkg/foo"
)

// "bar"
// "fmt"
// "net/http"
// "os"

// "github.com/go-chi/chi"

var Thingy string

func init() {
	Thingy = "Foo is initializing..."
	// fmt.Println(bar.ReturnBazHello())
}

func main() {
	// httpPort := os.Getenv("HTTP_PORT")
	// if httpPort == "" {
	// 	httpPort = "8080"
	// }
	// fmt.Println(Thingy)
	// http.HandleFunc("/foo", HelloFooService)
	// fmt.Println("Starting foo service...")
	// http.ListenAndServe(":"+httpPort, nil)

	fmt.Println(foo.UseFoo())
}

// func run() {
// if op == "http_test" {
// 	// mocking the pet svc to test our http routes
// 	svc := mocks.PetSvc{}
// 	h := ihttp.NewHandler(svc)
// 	r := chi.NewRouter()
// 	ihttp.Routes(r, h)
// 	return http.ListenAndServe(":6000", r)
// }
// }

// func HelloFooService(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.URL)
// 	fmt.Fprintf(w, "GET foo service endpoint: %s\n", r.URL.Path)
// }
