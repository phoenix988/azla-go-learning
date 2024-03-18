package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func mainScreenHandler(w http.ResponseWriter, data PageData) {

	// Parse the HTML template
	tmpl, err := template.ParseFiles("index/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template and write the result to the response
	tmpl.Execute(w, data)

}

func questionHandler(w http.ResponseWriter, r *http.Request) {
	data.CurrentQuestion = 0
	data.MaxAmountOfWords = 0
	data.CorrectAnswers = 0
	data.InCorrectAnswers = 0
	var words = []string{}
	var correct = []string{}

	var wordlist = create_wordlist()

	if r.FormValue("examMode") != "" {
		data.ExamMode = true
	} else {
		data.ExamMode = false
	}

	selectedWordList := r.FormValue("wordListOpt")
	selectedLanguage := r.FormValue("languageOpt")
	selectedCount := r.FormValue("wordCountOptions")
	fmt.Println(selectedCount)

	data.SelectedWordList = selectedWordList
	data.SelectedLanguage = selectedLanguage

	// Convert selected word count to int
	data.MaxAmountOfWords = convertToNum(selectedCount)

	// Create the word and correct slices/arrays
	for key, value := range wordlist[selectedWordList] {
		switch data.SelectedLanguage {
		case "Azerbajani":
			words = append(words, key)
			correct = append(correct, value)
		case "English":
			words = append(words, value)
			correct = append(correct, key)
		}
	}

	// Current word
	currentWord := words[data.CurrentQuestion]
	data.Words = words
	data.Correct = correct

	htmlStr := fmt.Sprintf("<form hx-post='/submit'>"+
		"<h1 class='app-title'>AZLA</h1><p class='word-list'>What is %s in %s ?<p>"+
		"<input type='text' name='answer'><button type='submit' name='evaluate'>Submit</button>"+
		"</form><form><button type='submit'>Back</button></form>", currentWord, selectedLanguage)

	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)

}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	currentCorrect := strings.ToLower(data.Correct[data.CurrentQuestion])
	data.CurrentCorrect = data.Correct[data.CurrentQuestion]
	userAnswer := strings.ToLower(r.FormValue("answer"))

	// Evaluates users response
	if userAnswer == currentCorrect {
		htmlStr := fmt.Sprintf("<form hx-post='/next'>"+
			"<h1 class='app-title'>AZLA</h1><p class='word-list'>Your answer is correct %s"+
			"<p><button type='submit' name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

		data.CorrectAnswers += 1
	} else {
		htmlStr := fmt.Sprintf("<form hx-post='/next'>"+
			"<h1 class='app-title'>AZLA</h1><p class='word-list'>Your answer is InCorrect %s<p>"+
			"<p>Correct answer is {{.CurrentCorrect}}</p><button type='submit' name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)
		data.InCorrectAnswers += 1
	}

	data.CurrentQuestion += 1

}

// Next question Handler
func nextHandler(w http.ResponseWriter, r *http.Request) {
	if data.CurrentQuestion >= len(data.Words) || data.CurrentQuestion >= data.MaxAmountOfWords {
		// data.CurrentQuestion = 0 // Reset to the first question if we reach the end
		// fmt.Println("Last question")
		htmlStr := fmt.Sprintf("<form hx-post='/'>" +
			"<h1 class='app-title'>AZLA</h1><p class='word-list'>You Have Reached The End<p>" +
			"<p>Number of Incorrect {{.InCorrectAnswers}}<p/><p>Number of Correct {{.CorrectAnswers}}</p>" +
			"</form>")

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)

	} else {
		currentWord := data.Words[data.CurrentQuestion]

		htmlStr := fmt.Sprintf("<form hx-post='/submit'>"+
			"<h1 class='app-title'>AZLA</h1><p class='word-list'>What is %s in %s ?<p>"+
			"<input type='text' name='answer'>"+
			"<button type='submit' name='evaluate'>Submit</button>"+
			"</form>", currentWord, data.SelectedLanguage)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

	}

}

// Main Handler
func handler(w http.ResponseWriter, r *http.Request) {
	// Create wordlist
	var wordlist = create_wordlist()
	var wordListOptions = []string{}

	// Define the Wordlist options and append them to the wordListOptions slice
	for key := range wordlist {
		wordListOptions = append(wordListOptions, key)

		//for i := 0; i < len(wordlist[key]); i++ {
		//	amountOfWords = append(amountOfWords, i+1)
		//}
	}

	// Create a PageData struct to pass to the template
	data := PageData{
		WordListOptions:         wordListOptions,
		LanguageOptions:         languageOptions,
		MaxAmountOfWordsOptions: amountOfWords,
	}

	mainScreenHandler(w, data)

}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/question", questionHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/next", nextHandler)
	http.ListenAndServe(":8080", nil)
}
