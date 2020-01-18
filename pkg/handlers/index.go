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

	//Display 404 if it's odball URL
	if r.URL.Path != "/" {
		page = template.Must(template.ParseFiles(
			"static/html/_base.html",
			"static/html/404.html",
		))

	}

	if r.Method == "GET" {
		page.Execute(w, params)
		//	return
	}

}
