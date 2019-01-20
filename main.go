package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

// HelloTemplateParams holds the data needed to render the hello page
type HelloTemplateParams struct {
	Name string
}

func writeTemplate(w io.Writer, path string, templateParams interface{}) error {
	// instantiate a template from the file
	t, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	// fill in the data from templateParams in the template and write the
	// result to the ResponseWriter
	if err := t.Execute(w, templateParams); err != nil {
		return err
	}
	return nil
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	// set up the parameters for the template
	templateParams := HelloTemplateParams{
		Name: r.URL.Path[1:],
	}
	err := writeTemplate(w, "templates/hello.html", templateParams)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func applyKernelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "Oh, hi")
	} else if r.Method == "POST" {
		fmt.Fprintln(w, "TBD")
	}
}

func main() {
	// Set handler functions
	http.HandleFunc("/", hiHandler)
	http.HandleFunc("/apply-kernel", applyKernelHandler)

	// Start the server and log any error it returns. The call to
	// http.ListenAndServe will only return when the server stops for some
	// reason.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
