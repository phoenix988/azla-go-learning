package main

// Static variables I set here
// Plus the main pagedata I want to pass to the template

import (
	"html/template"
)

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
	AvailableWords          []string // For all availble words 
	Correct                 []string // For all correct answers
	CurrentCorrect          string   // For the current correct answer
	CurrentWord             string   // For the current correct answer
	LanguageOptions         []string // Language Options
	CurrentQuestion         int      // Current question index
	CurrentIndex            int      // Current question index
	CorrectAnswers          int      // Number of correct answer
	InCorrectAnswers        int      // Number of incorrcet answers
	MaxAmountOfWords        int      // Max amount fo questions to ask
	MaxAmountOfWordsOptions []int    // Max amount fo questions to ask
	ExamMode                bool
	ExamModeAction          string
	ExamModeString          string
	IsComplete              map[string]bool
	CorrectAnswersList      map[string]string
	InCorrectAnswersList    map[string]string
}

// Create table for the data to pass to the template
var data = PageData{}

func create_questionString() (*template.Template, error) {
	tmpl, err := template.ParseFiles("index/questionAsk.html")

	return tmpl, err
}
