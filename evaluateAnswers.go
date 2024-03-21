package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func evaluateAnswers(userAnswer string, currentCorrect string, w http.ResponseWriter) {
	// Using switch statement to evaluate the answers
	switch userAnswer {
	case currentCorrect:
		htmlStr := fmt.Sprintf("<form hx-post='/next'>"+
			"<img class='image' src='https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg' alt='alternative-text' width='300' height='150'>"+
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
			"<img class='image' src='https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg' alt='alternative-text' width='300' height='150'>"+
			"<h1 class='app-title'>AZLA</h1><h2 class='word-list'>Your answer is Incorrect %s<p>"+
			"</h2><h2>Correct answer is: <h2 style='color: #ff5555;'>{{.CurrentCorrect}}</h2></h2><button type='submit'"+
			"name='next'>Next</button>"+
			"</form>", userAnswer)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)

		if data.IsComplete[currentCorrect] == false {
			data.InCorrectAnswers += 1
			data.IsComplete[currentCorrect] = true
			data.InCorrectAnswersList[currentCorrect + " Word: " + data.CurrentWord] = userAnswer
		}
	}

}
