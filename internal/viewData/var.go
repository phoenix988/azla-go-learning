package viewData

import (
	"github.com/gorilla/sessions"
	"html/template"
)

type PageData struct {
	WordListOptions         []string // Wordlist options
	WordList                map[string]map[string]string
	WordListName            string
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
	ExamModeAnswers         []string
	IsComplete              map[string]bool
	CorrectAnswersList      map[string]string
	InCorrectAnswersList    map[string]string
	CreateUser              bool
	IsSignedIn              bool
	FailedLoginAttempt      bool
	LoginUserName           string
	UserAnswer              string
	IsCorrect               bool
	WordImage               string
	CreateUserMes           string
}


var Data = PageData{}

// Language options
var LanguageOptions []string = []string{"Azerbajani", "English"}

// Set the wordlist count options
var AmountOfWords = []int{5, 10, 15, 20, 25, 30} // Amount of words to choose from

// Create session store
var Store = sessions.NewCookieStore([]byte("secret-key"))

func CreateQuestionTemp() (*template.Template, error) {
	tmpl, err := template.ParseFiles("index/questionAsk.html")

	return tmpl, err
}
