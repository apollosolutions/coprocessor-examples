package main

import (
	"github.com/rs/zerolog/log"
	"net/http"

	coprocessor "example.com/variables/internal"
)

func main() {

	http.HandleFunc("/", coprocessor.RequestHandler)

	log.Printf("Starting server on port 3007")
	http.ListenAndServe(":3007", nil)

}
