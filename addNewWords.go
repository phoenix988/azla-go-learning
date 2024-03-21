package main

import (
	"html/template"
	"net/http"
)

type dataToSave struct {
	WordlistName string
	IsWordlist   bool
	IsWordlistEmpty   bool // Check if you did enter a wordlist name or not
}

// Add new words to the app
func addWordListHandler(w http.ResponseWriter, r *http.Request) {

	DataToSave := dataToSave{}

	DataToSave.WordlistName = r.FormValue("addWordList")

	if r.FormValue("addWordList") != "" {
		DataToSave.IsWordlist = true
		DataToSave.IsWordlistEmpty = false
	} else { 
		DataToSave.IsWordlistEmpty = true
	}

	//.english := []string{"test", "test2", "test3", "test4"}
	//.aze := []string{"test", "test2", "test3", "test4"}

	//.saveWordJson(wordListName, english, aze)

	tmpl, _ := template.ParseFiles("index/addWords.html")
	tmpl.Execute(w, DataToSave)

}

func addWordsHandler(w http.ResponseWriter, r *http.Request) {

}

func addWordMainHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("index/addWords.html")

	tmpl.Execute(w, nil)

}
