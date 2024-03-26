package char

import (
	"strings"
    "net/http"
    "azla_go_learning/internal/viewData"
    "html/template"
)

func VariationLoop(variations []string, specialCharacters map[string]string) []string {
	for _, variation := range variations {
		for _, r := range variation {
			char := string(r)
			if replacement, ok := specialCharacters[char]; ok {
				newWord := strings.Replace(variation, char, replacement, 1)

				variations = append(variations, newWord)

			}
		}

	}

	return variations
}

func GenerateVariations(word string) []string {
	var specialCharacters = map[string]string{}
	specialCharacters = map[string]string{
		"ə": "e",
		"ü": "u",
		"ö": "o",
		"ı": "i",
		"ğ": "g",
		"ş": "s",
		"ç": "c",
		"i": "i",
	}

	var variations = []string{}
	variations = append(variations, word)

	variations = VariationLoop(variations, specialCharacters)
	variations = VariationLoop(variations, specialCharacters)
	variations = VariationLoop(variations, specialCharacters)


	return variations
}


func AppendStringAtIndex(slice []string, value string, index int) []string {
	// Ensure the index is within the bounds of the slice
	if index < 0 || index > len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}

	// Create a new slice with enough capacity to hold the additional element
	result := make([]string, len(slice)+1)

	// Copy the elements from the original slice up to the specified index
	copy(result, slice[:index])

	// Append the new value
	result[index] = value

	// Copy the remaining elements from the original slice
	copy(result[index+1:], slice[index:])

	return result
}



// Checks if you have minor mistakes
func EvaluateAnswers(userAnswer string, currentCorrect string, w http.ResponseWriter) {
	checkAltCorrect := false
	variations := GenerateVariations(currentCorrect)

	for _, altCorrect := range variations {
		if altCorrect == userAnswer {
			checkAltCorrect = true
			break
		}

	}

	// Using switch statement to evaluate the answers
	if userAnswer == currentCorrect || checkAltCorrect {
		viewData.Data.IsCorrect = true
		viewData.Data.UserAnswer = userAnswer

		tmpl, err := template.ParseFiles("index/response.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, viewData.Data)

		if viewData.Data.IsComplete[currentCorrect] == false {
			viewData.Data.CorrectAnswers += 1
			viewData.Data.IsComplete[currentCorrect] = true
			viewData.Data.CorrectAnswersList[currentCorrect] = userAnswer
		}
	} else {
		viewData.Data.IsCorrect = false
		viewData.Data.UserAnswer = userAnswer

		tmpl, err := template.ParseFiles("index/response.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, viewData.Data)

		if viewData.Data.IsComplete[currentCorrect] == false {
			viewData.Data.InCorrectAnswers += 1
			viewData.Data.IsComplete[currentCorrect] = true
			viewData.Data.InCorrectAnswersList[currentCorrect+" Word: "+viewData.Data.CurrentWord] = userAnswer
		}
	}

}
