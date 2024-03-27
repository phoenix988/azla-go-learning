package routes

import (
	"azla_go_learning/internal/char"
	"azla_go_learning/internal/json"
	"azla_go_learning/internal/viewData"
	"azla_go_learning/internal/words"
	"fmt"
	"html/template"
	"net/http"
)

// Parse index.html
func MainMenuIndex(w http.ResponseWriter, data viewData.PageData) {

	// Parse the HTML template
	tmpl, err := template.ParseFiles(viewData.TemplatePath + "menu.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template and write the result to the response
	tmpl.Execute(w, data)

}

func LoadCustomWords(wordlist map[string]map[string]string, r *http.Request) {
	session, _ := viewData.Store.Get(r, "session-name")
	username := session.Values["username"]

	importWordsFromJson, _ := jsonMod.ReadWordJson(jsonMod.JsonPath, username.(string))

	for key, value := range importWordsFromJson.Wordlist[username.(string)] {
		wordlist[key] = value
	}

	for key, value := range importWordsFromJson.Wordlist[username.(string)] {
		if _, ok := wordlist[key]; !ok {
			wordlist[key] = value
		}
	}

}

// Load main menu
func LoadMainMenu(w http.ResponseWriter, r *http.Request) {
	// Create wordlist
	var wordlist = words.CreateWordlist()

	// Retrive session
	session, _ := viewData.Store.Get(r, "session-name")

	// Create a PageData struct to pass to the template
	viewData.Data = viewData.PageData{
		LanguageOptions:         viewData.LanguageOptions, // language options
		MaxAmountOfWordsOptions: viewData.AmountOfWords,   // max count options
	}

	// Check if user is authenticated
	userID, ok := session.Values["user_id"].(int)
	if !ok || userID == 0 {
		// User is not authenticated, redirect to login page
		viewData.Data.IsSignedIn = false // not authenticated
		MainMenuIndex(w, viewData.Data)  // load login screen
		return
	} else {
		viewData.Data.IsSignedIn = true                 // authenitcated
		username := session.Values["username"]          // user
		viewData.Data.LoginUserName = username.(string) // convert to string

		// Read the json data file and append new words if they exist
		LoadCustomWords(wordlist, r)

		var wordListOptions = []string{}

		// Define the Wordlist options and append them to the wordListOptions slice
		for key := range wordlist {
			wordListOptions = append(wordListOptions, key)

			//for i := 0; i < len(wordlist[key]); i++ {
			//	amountOfWords = append(amountOfWords, i+1)
			//}
		}

		viewData.Data.WordListOptions = wordListOptions

		MainMenuIndex(w, viewData.Data)
		fmt.Println("User is authenticated")
	}

}

// Load the new questions
func LoadQuestions(w http.ResponseWriter, r *http.Request) {
	// Resets values
	viewData.Data.CurrentQuestion = 0
	viewData.Data.MaxAmountOfWords = 0
	viewData.Data.CorrectAnswers = 0
	viewData.Data.InCorrectAnswers = 0
	viewData.Data.IsComplete = map[string]bool{}
	viewData.Data.CorrectAnswersList = map[string]string{}
	viewData.Data.InCorrectAnswersList = map[string]string{}
	viewData.Data.AvailableWords = []string{}

	// Create wordlists
	wordlist := words.CreateWordlist()

	LoadCustomWords(wordlist, r) // Load custom wordlist if exist

	viewData.Data.WordList = wordlist // Assing wordlist to PageData

	// Arrays for words/correct answer
	var words = []string{}
	var correct = []string{}

	// Checks for exam mode
	if r.FormValue("examMode") != "" {
		viewData.Data.ExamMode = true
		viewData.Data.ExamModeAction = "/next" // Handler to use
		viewData.Data.ExamModeString = "Next"  // Label for button
	} else {
		viewData.Data.ExamMode = false
		viewData.Data.ExamModeAction = "/submit" // Handler to use
		viewData.Data.ExamModeString = "Submit"  // Label for button
	}

	selectedWordList := r.FormValue("wordListOpt")   // wordlist choice
	selectedLanguage := r.FormValue("languageOpt")   // Language choice
	selectedCount := r.FormValue("wordCountOptions") // Selected word count

	// Add to the PageData to pass to the templates
	viewData.Data.SelectedWordList = selectedWordList
	viewData.Data.SelectedLanguage = selectedLanguage

	// Convert selected word count to int
	viewData.Data.MaxAmountOfWords = char.ConvertToNum(selectedCount)
	viewData.Data.ExamModeAnswers = make([]string, viewData.Data.MaxAmountOfWords)

	// Create the word and correct slices/arrays
	for key, value := range viewData.Data.WordList[selectedWordList] {
		switch viewData.Data.SelectedLanguage {
		case "Azerbajani":
			words = append(words, key)
			correct = append(correct, value)
		case "English":
			words = append(words, value)
			correct = append(correct, key)
		}
	}

	// Current word and current index
	currentWord := words[viewData.Data.CurrentQuestion]
	currentIndex := viewData.Data.CurrentQuestion + 1

	// Search picture api for appropiate picture
	imageURLs, err := viewData.SearchImages(currentWord) // Example search query
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var selectedUrl string
	for _, url := range imageURLs {
		selectedUrl = url
		break
	}

	// Append to the PageData
	viewData.Data.Words = words
	viewData.Data.Correct = correct
	viewData.Data.WordImage = selectedUrl

	// Check if the list contains enough words from MaxAmountOfWords
	if viewData.Data.MaxAmountOfWords >= len(viewData.Data.Words) {
		fmt.Println("Wordlist doesn't have enough words:", len(viewData.Data.Words))
		fmt.Println("Choosen option:", viewData.Data.MaxAmountOfWords)
	}

	// Inserts zero for the web interface
	viewData.Data.AvailableWords = append(viewData.Data.AvailableWords, "")

	// Append all available words for the session
	for i := 0; i < len(viewData.Data.Words); i++ {
		if i >= viewData.Data.MaxAmountOfWords {
			break
		} else {
			viewData.Data.AvailableWords = append(viewData.Data.AvailableWords, viewData.Data.Words[i])
		}
	}

	// Print the available words to the terminal
	fmt.Println(viewData.Data.AvailableWords)

	if viewData.Data.ExamMode { // Checks for exam mode enable/disabled
		// Append to the PageDate
		viewData.Data.CurrentWord = currentWord
		viewData.Data.CurrentIndex = currentIndex
		tmpl, err := viewData.CreateQuestionTemp() // template from the main question

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, viewData.Data)

	} else {

		tmpl, err := viewData.CreateQuestionTemp() // template from the main question

		viewData.Data.CurrentWord = currentWord
		viewData.Data.CurrentIndex = currentIndex

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, viewData.Data)
	}

}
