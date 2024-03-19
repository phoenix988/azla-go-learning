package main

// Static variables I set here
// Plus the main pagedata I want to pass to the template

import "fmt"

// Language options
var languageOptions []string = []string{"Azerbajani", "English"}

// Set the wordlist count options
var amountOfWords = []int{5, 10, 15, 20, 25, 30} // Amount of words to choose from

// PageData struct to pass to the template
type PageData struct {
	WordListOptions         []string // Wordlist options
	SelectedWordList        string   // Selected wordlist option
	SelectedLanguage        string   // Selected Language Option
	Words                   []string // For all words
	Correct                 []string // For all correct answers
	CurrentCorrect          string   // For the current correct answer
	CurrentWord             string   // For the current correct answer
	LanguageOptions         []string // Language Options
	CurrentQuestion         int      // Current question index
	CorrectAnswers          int      // Number of correct answer
	InCorrectAnswers        int      // Number of incorrcet answers
	MaxAmountOfWords        int      // Max amount fo questions to ask
	MaxAmountOfWordsOptions []int    // Max amount fo questions to ask
	ExamMode                bool
}

// Create table for the data to pass to the template
var data = PageData{}

func create_questionString(currentIndex int, currentWord string, data PageData, mode string)string {
	var buttonType string
	var buttonName string
	
	switch mode {
	case "/submit":
		buttonType = "Submit"
		buttonName = "evaluate"
		
	case "/next":
		buttonType = "Next"
		buttonName = "next"

	}
	var questionString = fmt.Sprintf("<form hx-post='%s'>"+
		"<h1 class='app-title'>AZLA</h1><p class='word-list'>%d What is <span class='wordQuestion'>%s </span>" + 
		"in<span class='wordLanguage'> %s </span>?<p>"+
		"<input type='text' name='answer'>"+
		"<button type='submit' name='%s'>%s</button>"+
		"</form>",mode, currentIndex, currentWord, data.SelectedLanguage, buttonName, buttonType)


		return questionString
}
