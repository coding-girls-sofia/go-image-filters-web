package main

import (
	"html/template"
	"log"
	"net/http"
)

// HelloTemplateParams holds the data needed to render the hello page
type HelloTemplateParams struct {
	Name string
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	// set up the parameters for the template
	templateParams := HelloTemplateParams{
		Name: r.URL.Path[1:],
	}
	// instantiate a template from the file
	t, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// fill in the data from templateParams in the template and write the
	// result to the ResponseWriter
	if err := t.Execute(w, templateParams); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	// Set the handler function for the root URI
	http.HandleFunc("/", hiHandler)

	// Start the server and log any error it returns. The call to
	// http.ListenAndServe will only return when the server stops for some
	// reason.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
