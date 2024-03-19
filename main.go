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
			"<img class='image' src='https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg' alt='alternative-text' width='300' height='150'>" +
			"<h1 class='app-title'>AZLA</h1><h2 class='word-list'>Your answer is correct %s"+
			"</h2><button type='submit' name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

		if data.IsComplete[currentCorrect] == false {
			data.CorrectAnswers += 1
			data.IsComplete[currentCorrect] = true
			data.CorrectAnswersList[currentCorrect] = userAnswer
		}

	default:
		htmlStr := fmt.Sprintf("<form hx-post='/next'>"+
			"<img class='image' src='https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg' alt='alternative-text' width='300' height='150'>" +
			"<h1 class='app-title'>AZLA</h1><h2 class='word-list'>Your answer is Incorrect %s<p>"+
			"</h2><h2>Correct answer is: <h2 style='color: #ff5555;'>{{.CurrentCorrect}}</h2></h2><button type='submit'"+
			"name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)
		
		if data.IsComplete[currentCorrect] == false {
			data.InCorrectAnswers += 1
			data.IsComplete[currentCorrect] = true
			data.InCorrectAnswersList[currentCorrect] = userAnswer
		}
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
	data.IsComplete = map[string]bool{}
	data.CorrectAnswersList = map[string]string{}
	data.InCorrectAnswersList = map[string]string{}
	data.AvailableWords = []string{}
	var words = []string{}
	var correct = []string{}

	var wordlist = create_wordlist()

	if r.FormValue("examMode") != "" {
		data.ExamMode = true
		data.ExamModeAction = "/next"
		data.ExamModeString = "Next"
	} else {
		data.ExamMode = false
		data.ExamModeAction = "/submit"
		data.ExamModeString = "Submit"
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
		fmt.Println("Wordlist doesn't have enough words:", len(data.Words))
		fmt.Println("Choosen option:", data.MaxAmountOfWords)
	}

	for i := 0; i < len(data.Words); i++ {
		if i > data.MaxAmountOfWords { 
			break
			} else {
			fmt.Println("runs")
			data.AvailableWords = append(data.AvailableWords, data.Words[i])
		}
	}

	fmt.Println(data.AvailableWords)


	if data.ExamMode {

		data.CurrentWord = currentWord
		data.CurrentIndex = currentIndex
		tmpl, _ := create_questionString()

		//tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)

	} else {

		//htmlStr := create_questionString(currentIndex, currentWord, data, "/submit")
		tmpl, err := create_questionString()

		//tmpl, _ := template.New("t").Parse(htmlStr)
		data.CurrentWord = currentWord
		data.CurrentIndex = currentIndex

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
		
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
		if data.ExamMode { // If exam mode is activated
			currentCorrect := strings.ToLower(data.Correct[data.CurrentQuestion])
			userAnswer := strings.ToLower(r.FormValue("answer"))
			switch userAnswer {
			case currentCorrect:
				if data.IsComplete[currentCorrect] == false {
					data.CorrectAnswers += 1
					data.IsComplete[currentCorrect] = true
					data.CorrectAnswersList[currentCorrect] = userAnswer
					fmt.Println("Correct")
				}
			default:
				if data.IsComplete[currentCorrect] == false {
					data.InCorrectAnswers += 1
					data.IsComplete[currentCorrect] = true
					data.InCorrectAnswersList[currentCorrect] = userAnswer
					fmt.Println("Incorrect, Correct:", data.CurrentCorrect)
				}
			}
		}

		fmt.Println("correct: ", data.CorrectAnswers)
		fmt.Println("incorrect: ", data.InCorrectAnswers)
		
		// parse the end menu html file that displays the results
		tmpl, _ := template.ParseFiles("index/endMenu.html")

		tmpl.Execute(w, data)

	} else {
		data.CurrentQuestion += 1
		if data.ExamMode {

			currentCorrect := strings.ToLower(data.Correct[data.CurrentQuestion-1])
			userAnswer := strings.ToLower(r.FormValue("answer"))

			switch userAnswer {
			case currentCorrect:
				if data.IsComplete[currentCorrect] == false {
					data.CorrectAnswers += 1
					data.IsComplete[currentCorrect] = true
					data.CorrectAnswersList[currentCorrect] = userAnswer
				}
				fmt.Println("Correct")
			default:
				if data.IsComplete[currentCorrect] == false {
					data.InCorrectAnswers += 1
					data.IsComplete[currentCorrect] = true
					data.InCorrectAnswersList[currentCorrect] = userAnswer
				}
				fmt.Println("Incorrect, Correct:", currentCorrect)
			}

			currentWord := data.Words[data.CurrentQuestion]
			currentIndex := data.CurrentQuestion + 1
			data.CurrentWord = currentWord
			data.CurrentIndex = currentIndex

			htmlStr, _ := create_questionString()

			//tmpl, _ := template.New("t").Parse(htmlStr)
			htmlStr.Execute(w, data)
		} else {
			currentWord := data.Words[data.CurrentQuestion]
			currentIndex := data.CurrentQuestion + 1
			data.CurrentWord = currentWord
			data.CurrentIndex = currentIndex

			htmlStr, _ := create_questionString()

			htmlStr.Execute(w, data)

		}

	}

}

func prevHandler(w http.ResponseWriter, r *http.Request) {

	currentWord := data.Words[data.CurrentQuestion-1]
	currentIndex := data.CurrentQuestion
	fmt.Println(data.CurrentQuestion)
	var htmlStr string
	htmlStr = fmt.Sprintf("<p id='wordQuestion'>" +
		"<span id='wordQuestion'>%d </span>What is " +
		"<span id='wordQuestion' class='wordQuestion'>%s" +
		"</span> in <span class='wordLanguage'> %s</span> ?</p>", currentIndex, currentWord, data.SelectedLanguage)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occures")
	}
	tmpl.Execute(w, nil)

	data.CurrentQuestion -= 1

}

func wordlistChangeHandler(w http.ResponseWriter, r *http.Request) {
	selectedWordList := r.FormValue("wordListOpt")
	htmlStr := fmt.Sprintf("<div id='wordListTitle'>Wordlist <span class='wordQuestion'>%s</span> Selected</div>", selectedWordList)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	tmpl.Execute(w, nil)

}

func wordResultHandler(w http.ResponseWriter, r *http.Request) {

	for key, value := range data.InCorrectAnswersList {
		fmt.Println(key, value)
	}

	tmpl, _ := template.ParseFiles("index/resultMenu.html")

	tmpl.Execute(w, data)
}


func wordcountChangeHandler(w http.ResponseWriter, r *http.Request) {
	selectedCount := r.FormValue("wordCountOptions")
	htmlStr := fmt.Sprintf("<div id='wordAmount'><span class='wordQuestion'>%s</span> Words Selected</div>", selectedCount)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	tmpl.Execute(w, nil)

}

func languageChangeHandler(w http.ResponseWriter, r *http.Request) {
	selectedLanguage := r.FormValue("languageOpt")
	htmlStr := fmt.Sprintf("<div id='languageTitle'><span class='wordQuestion'>%s</span> Selected</div>", selectedLanguage)

	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	tmpl.Execute(w, nil)

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
	http.HandleFunc("/wordcount_changed", wordcountChangeHandler)
	http.HandleFunc("/language_changed", languageChangeHandler)
	http.HandleFunc("/word_result", wordResultHandler)
	http.ListenAndServe(":8080", nil)
}
