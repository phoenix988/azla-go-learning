package routes

import (
	"html/template"
	"fmt"
	"net/http"

)


func wordlistChangeHandler(w http.ResponseWriter, r *http.Request) {
	selectedWordList := r.FormValue("wordListOpt")
	htmlStr := fmt.Sprintf("<div id='wordListTitle'>Wordlist <span class='wordQuestion'>%s</span> Selected</div>", selectedWordList)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	tmpl.Execute(w, nil)

}

func wordCountChangeHandler(w http.ResponseWriter, r *http.Request) {
	selectedCount := r.FormValue("wordCountOptions")
	htmlStr := fmt.Sprintf("<div id='wordAmount'><span class='wordQuestion'>%s</span> Words Selected</div>", selectedCount)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	tmpl.Execute(w, nil)

}

func languageChangeHandler(w http.ResponseWriter, r *http.Request) {
	selectedLanguage := r.FormValue("languageOpt")
	htmlStr := fmt.Sprintf("<div id='languageTitle'><span class='wordQuestion'>%s</span> Selected</div>", selectedLanguage)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	tmpl.Execute(w, nil)

}

func NewUpdateStringRoutes() {
	http.HandleFunc("/wordlist_changed", wordlistChangeHandler)
	http.HandleFunc("/wordcount_changed", wordCountChangeHandler)
	http.HandleFunc("/language_changed", languageChangeHandler)

}
