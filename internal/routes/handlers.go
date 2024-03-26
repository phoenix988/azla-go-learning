package routes

import (
	"azla_go_learning/internal/char"
	"azla_go_learning/internal/json"
	"azla_go_learning/internal/viewData"
	"azla_go_learning/internal/words"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)



func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Create wordlist
	var wordlist = words.CreateWordlist()

	session, _ := viewData.Store.Get(r, "session-name")

	// Create a PageData struct to pass to the template
	viewData.Data = viewData.PageData{
		LanguageOptions:         viewData.LanguageOptions,
		MaxAmountOfWordsOptions: viewData.AmountOfWords,
	}

	// Check if user is authenticated
	userID, ok := session.Values["user_id"].(int)
	if !ok || userID == 0 {
		// User is not authenticated, redirect to login page
		viewData.Data.IsSignedIn = false
		MainMenuIndex(w, viewData.Data)
		return
	} else {
		viewData.Data.IsSignedIn = true
		username := session.Values["username"]
		viewData.Data.LoginUserName = username.(string)
		// Read the json data file and append new words if they exist
		importWordsFromJson, _ := jsonMod.ReadWordJson(jsonMod.JsonPath, viewData.Data.LoginUserName)

		for key, value := range importWordsFromJson.Wordlist[viewData.Data.LoginUserName] {
			wordlist[key] = value
		}

		for key, value := range importWordsFromJson.Wordlist[viewData.Data.LoginUserName] {
			if _, ok := wordlist[key]; !ok {
				wordlist[key] = value
			}
		}

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

// Next question Handler
func NextHandler(w http.ResponseWriter, r *http.Request) {
	// Check if you are on the last question
	if viewData.Data.CurrentQuestion >= len(viewData.Data.Words) || viewData.Data.CurrentQuestion >= viewData.Data.MaxAmountOfWords-1 {
		if viewData.Data.ExamMode { // If exam mode is activated
			//currentCorrect := strings.ToLower(viewData.Data.Correct[viewData.Data.CurrentQuestion])
			userAnswer := strings.ToLower(r.FormValue("answer"))
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

		}

		fmt.Println("correct: ", viewData.Data.CorrectAnswers)
		fmt.Println("incorrect: ", viewData.Data.InCorrectAnswers)

		// parse the end menu html file that displays the results
		tmpl, _ := template.ParseFiles("index/endMenu.html")

		tmpl.Execute(w, viewData.Data)

	} else {
		viewData.Data.CurrentQuestion += 1
		if viewData.Data.ExamMode {
			userAnswer := strings.ToLower(r.FormValue("answer"))

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
	userAnswer := strings.ToLower(r.FormValue("answer"))

	// Evaluates users response
	char.EvaluateAnswers(userAnswer, currentCorrect, w)

}
