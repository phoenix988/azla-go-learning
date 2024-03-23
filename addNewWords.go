package main

import (
	"html/template"
	"net/http"
	"azla_go_learning/internal/json"
	"fmt"
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
	data.WordListName = DataToSave.WordlistName 

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

	english := r.FormValue("engText")
	aze := r.FormValue("azeText")

	var englishWords []string
	var azeWords []string

	englishWords = append(englishWords, english)
	azeWords = append(azeWords, aze)

	session, _ := store.Get(r, "session-name")

	jsonMod.SaveWordJson(data.WordListName, englishWords, azeWords, session.Values["username"].(string))

}

func addWordMainHandler(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		username := session.Values["username"]
		data.LoginUserName = username.(string)
		// Read the json data file and append new words if they exist
		importWordsFromJson, _ := jsonMod.ReadWordJson(jsonMod.JsonPath, data.LoginUserName)

		for key, value := range importWordsFromJson.Wordlist[data.LoginUserName] {
			fmt.Println(key, value)
		}


	tmpl, _ := template.ParseFiles("index/addWords.html")

	tmpl.Execute(w, nil)

}
