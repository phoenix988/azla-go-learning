package routes

import (
	"azla_go_learning/internal/viewData"
	"html/template"
	"net/http"
)

// Parse index.html
func MainMenuIndex(w http.ResponseWriter, data viewData.PageData) {

	// Parse the HTML template
	tmpl, err := template.ParseFiles("index/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template and write the result to the response
	tmpl.Execute(w, data)

}

