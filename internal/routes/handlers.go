package routes

import (
	"azla_go_learning/internal/char"
	"azla_go_learning/internal/viewData"
	"azla_go_learning/internal/quiz"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)


func MainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		LoadMainMenu(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// when you first start the quiz
func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	LoadQuestions(w, r)
}

// Next question Handler
func NextHandler(w http.ResponseWriter, r *http.Request) {
	// Check if you are on the last question
	if viewData.Data.CurrentQuestion >= len(viewData.Data.Words) || viewData.Data.CurrentQuestion >= viewData.Data.MaxAmountOfWords-1 {
		quiz.IfLastQuestion(w, r) // if question is the last one it shows the end screen
	} else {
		quiz.NextQuestion(w, r) // Else it goes to the next question
	}

}

func PrevHandler(w http.ResponseWriter, r *http.Request) {

	currentWord := viewData.Data.Words[viewData.Data.CurrentQuestion-1]
	currentIndex := viewData.Data.CurrentQuestion

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

	viewData.Data.WordImage = selectedUrl
	viewData.Data.CurrentQuestion -= 1

	viewData.Data.CurrentWord = currentWord
	viewData.Data.CurrentIndex = currentIndex
	tmpl, _ := viewData.CreateQuestionTemp()

	tmpl.Execute(w, viewData.Data)

}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	currentCorrect := strings.ToLower(viewData.Data.Correct[viewData.Data.CurrentQuestion])
	viewData.Data.CurrentCorrect = viewData.Data.Correct[viewData.Data.CurrentQuestion]
	userAnswer := strings.ToLower(strings.TrimSpace(r.FormValue("answer")))

	// Evaluates users response
	char.EvaluateAnswers(userAnswer, currentCorrect, w)

}

func JumpHandler(w http.ResponseWriter, r *http.Request) {

	jumpTo := (r.FormValue("index"))

	newIndex := char.ConvertToNum(jumpTo)

	currentWord := viewData.Data.AvailableWords[newIndex]
	currentIndex := newIndex + 1

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

	//userAnswer := strings.ToLower(r.FormValue("answer"))
	//viewData.Data.ExamModeAnswers = char.AppendStringAtIndex(viewData.Data.ExamModeAnswers, userAnswer, viewData.Data.CurrentQuestion-1)
	viewData.Data.WordImage = selectedUrl
	viewData.Data.CurrentQuestion = newIndex

	viewData.Data.CurrentWord = currentWord
	viewData.Data.CurrentIndex = currentIndex
	viewData.Data.CurrentIndex -= 1
	//viewData.Data.CurrentQuestion = currentIndex-1

	tmpl, _ := viewData.CreateQuestionTemp()

	tmpl.Execute(w, viewData.Data)

	viewData.Data.CurrentQuestion -= 1

}

func WordResultHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range viewData.Data.InCorrectAnswersList {
		fmt.Println(key, value)
	}

	tmpl, _ := template.ParseFiles(viewData.TemplatePath + "resultMenu.html")

	tmpl.Execute(w, viewData.Data)
}

// Routes related to jumping between questions in the quiz
func NewMainRoutes() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/question", QuestionHandler)
	http.HandleFunc("/submit", SubmitHandler)
	http.HandleFunc("/next", NextHandler)
	http.HandleFunc("/prev", PrevHandler)
	http.HandleFunc("/jump", JumpHandler)
	http.HandleFunc("/word_result", WordResultHandler)
}
