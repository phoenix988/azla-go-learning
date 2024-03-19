package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func evaluateAnswers(userAnswer string, currentCorrect string, w http.ResponseWriter) {
	// Using switch statement to evaluate the answers
	switch userAnswer {
	case currentCorrect:
		htmlStr := fmt.Sprintf("<form hx-post='/next'>"+
			"<h1 class='app-title'>AZLA</h1><h2 class='word-list'>Your answer is correct %s"+
			"</h2><button type='submit' name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

		data.CorrectAnswers += 1
	default:
		htmlStr := fmt.Sprintf("<form hx-post='/next'>"+
			"<h1 class='app-title'>AZLA</h1><h2 class='word-list'>Your answer is Incorrect %s<p>"+
			"</h2><h2>Correct answer is: <h2 style='color: #ff5555;'>{{.CurrentCorrect}}</h2></h2><button type='submit'" +
			"name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)
		data.InCorrectAnswers += 1
	}

}

// Parse index.html
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
	currentIndex := data.CurrentQuestion + 1
	data.Words = words
	data.Correct = correct

	if data.MaxAmountOfWords >= len(data.Words) {
		fmt.Println("Wordlist doesnt cointain enough words", len(data.Words))

	}

	if data.ExamMode {
		htmlStr := create_questionString(currentIndex, currentWord, data, "/next")

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)

	} else {

		htmlStr := create_questionString(currentIndex, currentWord, data, "/submit")

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

	}

}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	currentCorrect := strings.ToLower(data.Correct[data.CurrentQuestion])
	data.CurrentCorrect = data.Correct[data.CurrentQuestion]
	userAnswer := strings.ToLower(r.FormValue("answer"))

	// Evaluates users response
	evaluateAnswers(userAnswer, currentCorrect, w)

}

// Next question Handler
func nextHandler(w http.ResponseWriter, r *http.Request) {
	// Check if you are on the last question
	if data.CurrentQuestion >= len(data.Words) || data.CurrentQuestion >= data.MaxAmountOfWords-1 {
		if data.ExamMode  { // If exam mode is activated
			currentCorrect := strings.ToLower(data.Correct[data.CurrentQuestion-1])
			userAnswer := strings.ToLower(r.FormValue("answer"))
			switch userAnswer {
			case currentCorrect:
				data.CorrectAnswers += 1
				fmt.Println("Correct")
			default:
				data.InCorrectAnswers += 1
				fmt.Println("Incorrect, Correct:", data.CurrentCorrect)
			}
		}

		htmlStr := fmt.Sprintf("<form hx-post='/'>" +
			"<h1 class='app-title'>AZLA</h1><p class='word-list'>You Have Reached The End<p>" +
			"<p>Number of Incorrect {{.InCorrectAnswers}}<p/><p>Number of Correct {{.CorrectAnswers}}</p>" +
			"<button type='submit'>Back</form>")

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)

	} else {
		data.CurrentQuestion += 1
		if data.ExamMode {

			currentCorrect := strings.ToLower(data.Correct[data.CurrentQuestion-1])
			userAnswer := strings.ToLower(r.FormValue("answer"))

			switch userAnswer {
			case currentCorrect:
				data.CorrectAnswers += 1
				fmt.Println("Correct")
			default:
				data.InCorrectAnswers += 1
				fmt.Println("Incorrect, Correct:", currentCorrect)
			}

			currentWord := data.Words[data.CurrentQuestion]
			currentIndex := data.CurrentQuestion + 1

			htmlStr := create_questionString(currentIndex, currentWord, data, "/next")
			
			tmpl, _ := template.New("t").Parse(htmlStr)
			tmpl.Execute(w, nil)
		} else {
			currentWord := data.Words[data.CurrentQuestion]
			currentIndex := data.CurrentQuestion + 1

			htmlStr := create_questionString(currentIndex, currentWord, data, "/submit")

			tmpl, _ := template.New("t").Parse(htmlStr)
			tmpl.Execute(w, nil)

		}

	}

}

func prevHandler(w http.ResponseWriter, r *http.Request) {



}


func wordlistChangeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")

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
	http.HandleFunc("/prev", prevHandler)
	http.HandleFunc("/wordlist_changed", wordlistChangeHandler)
	http.ListenAndServe(":8080", nil)
}
