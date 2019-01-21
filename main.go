package main

import (
	"errors"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"image/png"
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

func writeImage(w io.Writer, imageData image.Image, format string) error {
	switch format {
	case "jpeg":
		return jpeg.Encode(w, imageData, nil)
	case "png":
		return png.Encode(w, imageData)
	default:
		return errors.New("Unknown format")
	}
}

func applyKernelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		writeTemplate(w, "templates/apply-kernel.html", nil)
	} else if r.Method == "POST" {
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, fmt.Sprintf("reading file filed: %s", err.Error()), 400)
			return
		}
		defer file.Close()
		w.Header().Set("Content-Type", handler.Header.Get("Content-Type"))
		io.Copy(w, file)
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
