package routes

import (
	"encoding/json"
	"net/http"
	"azla_go_learning/internal/words"
)

func WordsApiHandler(w http.ResponseWriter, r *http.Request){
		var wordlist = words.CreateWordlist()
		wordlistJSON, err := json.Marshal(wordlist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		// Write JSON response
		w.Write(wordlistJSON)

}
