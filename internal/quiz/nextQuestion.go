package quiz

import (
	"azla_go_learning/internal/char"
	"azla_go_learning/internal/viewData"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func IfLastQuestion(w http.ResponseWriter, r *http.Request) {
	//currentCorrect := strings.ToLower(viewData.Data.Correct[viewData.Data.CurrentQuestion])
	userAnswer := strings.ToLower(strings.TrimSpace(r.FormValue("answer")))
	viewData.Data.ExamModeAnswers = char.AppendStringAtIndex(viewData.Data.ExamModeAnswers, userAnswer, viewData.Data.CurrentQuestion)
	for index := range viewData.Data.Correct {
		checkAltCorrect := false

		currentCorrect := strings.ToLower(viewData.Data.Correct[index])
		variations := char.GenerateVariations(currentCorrect)

		for _, altCorrect := range variations {
			if altCorrect == viewData.Data.ExamModeAnswers[index] {
				checkAltCorrect = true
				break
			}

		}

		if viewData.Data.ExamModeAnswers[index] == currentCorrect || checkAltCorrect {
			if viewData.Data.IsComplete[currentCorrect] == false {
				viewData.Data.CorrectAnswers += 1
				viewData.Data.IsComplete[currentCorrect] = true
				viewData.Data.CorrectAnswersList[currentCorrect] = viewData.Data.ExamModeAnswers[index]
			}
		} else {
			if viewData.Data.IsComplete[currentCorrect] == false {
				viewData.Data.InCorrectAnswers += 1
				viewData.Data.IsComplete[currentCorrect] = true
				viewData.Data.InCorrectAnswersList[currentCorrect+" Word: "+viewData.Data.Words[index]] = viewData.Data.ExamModeAnswers[index]
			}

		}
		if index >= len(viewData.Data.Words) || index >= viewData.Data.MaxAmountOfWords-1 {
			fmt.Println(index)
			break
		}

	}

	fmt.Println("correct: ", viewData.Data.CorrectAnswers)
	fmt.Println("incorrect: ", viewData.Data.InCorrectAnswers)

	// parse the end menu html file that displays the results
	tmpl, _ := template.ParseFiles(viewData.TemplatePath + "endMenu.html")

	tmpl.Execute(w, viewData.Data)

}

func NextQuestion(w http.ResponseWriter, r *http.Request) {
	viewData.Data.CurrentQuestion += 1
	if viewData.Data.ExamMode {
		userAnswer := strings.ToLower(strings.TrimSpace(r.FormValue("answer")))

		viewData.Data.ExamModeAnswers = char.AppendStringAtIndex(viewData.Data.ExamModeAnswers, userAnswer, viewData.Data.CurrentQuestion-1)

		currentWord := viewData.Data.Words[viewData.Data.CurrentQuestion]
		currentIndex := viewData.Data.CurrentQuestion + 1
		viewData.Data.CurrentWord = currentWord
		viewData.Data.CurrentIndex = currentIndex

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

		htmlStr, _ := viewData.CreateQuestionTemp()

		htmlStr.Execute(w, viewData.Data)
	} else {
		currentWord := viewData.Data.Words[viewData.Data.CurrentQuestion]
		currentIndex := viewData.Data.CurrentQuestion + 1
		viewData.Data.CurrentWord = currentWord
		viewData.Data.CurrentIndex = currentIndex

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

		htmlStr, _ := viewData.CreateQuestionTemp()

		htmlStr.Execute(w, viewData.Data)
	}
}
