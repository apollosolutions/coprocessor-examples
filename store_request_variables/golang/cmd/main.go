package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"

	coprocessor "example.com/variables/internal"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", coprocessor.RequestHandler)

	log.Printf("Starting server on port 3007")
	http.ListenAndServe(":3007", nil)

}
