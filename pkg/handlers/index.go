package handlers

import (
	"html/template"
	"net/http"
)

// TemplateParams is TemplateParams
type TemplateParams struct {
	// Notice    string
	// ID        int
	// FirstName string
	// LastName  string
	// EmailAddr string
	// Phone     string
	// Message   string
}

// Index is index.
func Index(w http.ResponseWriter, r *http.Request) {
	params := TemplateParams{}

	// Set the default page

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/index.html",
	))

	// TODO: Catch bad request methods, write appropriate responses

	if r.Method == "GET" {
		page.Execute(w, params)
		//	return
	}

}
