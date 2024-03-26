package main

import (
	"azla_go_learning/internal/json"
	"azla_go_learning/internal/viewData"
	"html/template"
	"net/http"
)

type dataToSave struct {
	WordlistName       string
	IsWordlist         bool
	IsWordlistEmpty    bool                           // Check if you did enter a wordlist name or not
	CustomWordlist     map[string]map[string]string   // Check if you did enter a wordlist name or not
	CustomWordlistName []map[string]map[string]string // Check if you did enter a wordlist name or not
}

// Add new words to the app
func addWordListHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")

	userID, ok := session.Values["user_id"].(int)
	if !ok || userID == 0 {

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		DataToSave := dataToSave{}

		DataToSave.WordlistName = r.FormValue("addWordList")
		viewData.Data.WordListName = DataToSave.WordlistName

		if r.FormValue("addWordList") != "" {
			DataToSave.IsWordlist = true
			DataToSave.IsWordlistEmpty = false
		} else {
			DataToSave.IsWordlistEmpty = true
		}

		tmpl, _ := template.ParseFiles("index/addWords.html")
		tmpl.Execute(w, DataToSave)

	}

}

func addWordsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	userID, ok := session.Values["user_id"].(int)
	if !ok || userID == 0 {

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		english := r.FormValue("engText")
		aze := r.FormValue("azeText")

		var englishWords []string
		var azeWords []string

		englishWords = append(englishWords, english)
		azeWords = append(azeWords, aze)

		session, _ := store.Get(r, "session-name")

		jsonMod.SaveWordJson(viewData.Data.WordListName, englishWords, azeWords, session.Values["username"].(string))

		http.Redirect(w, r, "/", http.StatusFound)

	}

}

func addWordMainHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	userID, ok := session.Values["user_id"].(int)
	if !ok || userID == 0 {

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		username := session.Values["username"]
		viewData.Data.LoginUserName = username.(string)
		// Read the json viewData.Data file and append new words if they exist
		importWordsFromJson, _ := jsonMod.ReadWordJson(jsonMod.JsonPath, viewData.Data.LoginUserName)

		DataToSave := dataToSave{
			CustomWordlist:     map[string]map[string]string{},
			CustomWordlistName: []map[string]map[string]string{}, // Check if you did enter a wordlist name or not
		}

		for key, value := range importWordsFromJson.Wordlist[viewData.Data.LoginUserName] {
			DataToSave.CustomWordlist[key] = value
			DataToSave.CustomWordlistName = append(DataToSave.CustomWordlistName, map[string]map[string]string{key: value})
		}

		tmpl, _ := template.ParseFiles("index/addWords.html")

		tmpl.Execute(w, DataToSave)

	}

}
