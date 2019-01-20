package main

import (
	"fmt"
	"log"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	// Write directly to the ResponseWriter
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func main() {
	// Set the handler function for the root URI
	http.HandleFunc("/", hiHandler)

	// Start the server and log any error it returns. The call to
	// http.ListenAndServe will only return when the server stops for some
	// reason.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
