package main

import (
	"azla_go_learning/internal/json"
	"azla_go_learning/internal/userAuth"
	"azla_go_learning/internal/words"
	//"azla_go_learning/internal/database"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// Parse index.html
func MainMenuIndex(w http.ResponseWriter, data PageData) {

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

	wordlist := words.CreateWordlist()

	session, _ := store.Get(r, "session-name")
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

	data.WordList = wordlist

	var words = []string{}
	var correct = []string{}

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
	for key, value := range data.WordList[selectedWordList] {
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
			data.AvailableWords = append(data.AvailableWords, data.Words[i])
		}
	}

	fmt.Println(data.AvailableWords)

	if data.ExamMode {

		data.CurrentWord = currentWord
		data.CurrentIndex = currentIndex
		tmpl, _ := CreateQuestionTemp()

		//tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, data)

	} else {

		//htmlStr := CreateQuestionTemp(currentIndex, currentWord, data, "/submit")
		tmpl, err := CreateQuestionTemp()

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
					data.InCorrectAnswersList[currentCorrect+" Word: "+data.CurrentWord] = userAnswer
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
					data.InCorrectAnswersList[currentCorrect+" Word: "+data.CurrentWord] = userAnswer
				}
				fmt.Println("Incorrect, Correct:", currentCorrect)
			}

			currentWord := data.Words[data.CurrentQuestion]
			currentIndex := data.CurrentQuestion + 1
			data.CurrentWord = currentWord
			data.CurrentIndex = currentIndex

			htmlStr, _ := CreateQuestionTemp()

			//tmpl, _ := template.New("t").Parse(htmlStr)
			htmlStr.Execute(w, data)
		} else {
			currentWord := data.Words[data.CurrentQuestion]
			currentIndex := data.CurrentQuestion + 1
			data.CurrentWord = currentWord
			data.CurrentIndex = currentIndex

			htmlStr, _ := CreateQuestionTemp()

			htmlStr.Execute(w, data)

		}

	}

}

func prevHandler(w http.ResponseWriter, r *http.Request) {

	currentWord := data.Words[data.CurrentQuestion-1]
	currentIndex := data.CurrentQuestion
	fmt.Println(data.CurrentQuestion)
	var htmlStr string
	htmlStr = fmt.Sprintf("<p class='questionString' id='wordQuestion'>"+
		"<span id='wordQuestion'>%d </span>What is "+
		"<span id='wordQuestion' class='wordQuestion'>%s"+
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
	var wordlist = words.CreateWordlist()

	session, _ := store.Get(r, "session-name")

	// Create a PageData struct to pass to the template
	data := PageData{
		LanguageOptions:         languageOptions,
		MaxAmountOfWordsOptions: amountOfWords,
	}
	

	// Check if user is authenticated
	userID, ok := session.Values["user_id"].(int)
	if !ok || userID == 0 {
		// User is not authenticated, redirect to login page
		data.IsSignedIn = false
		MainMenuIndex(w, data)
		return
	} else {

		data.IsSignedIn = true
		username := session.Values["username"]
		data.LoginUserName = username.(string)
		// Read the json data file and append new words if they exist
		importWordsFromJson, _ := jsonMod.ReadWordJson(jsonMod.JsonPath, data.LoginUserName)

		for key, value := range importWordsFromJson.Wordlist[data.LoginUserName] {
			wordlist[key] = value
		}

		for key, value := range importWordsFromJson.Wordlist[data.LoginUserName] {
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

		data.WordListOptions = wordListOptions

		MainMenuIndex(w, data)
		fmt.Println("User is authenticated")

	}

}

// Launch the login screen
func loginHandler(w http.ResponseWriter, r *http.Request) {
	data.CreateUser = false

	//	MainMenuIndex(w, data)

	// tmpl, _ := template.ParseFiles("index/signIn.html")

	// tmpl.Execute(w, data)

}

// Handle logout event
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Revoke authentication
	delete(session.Values, "user_id")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

// Create the user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	data.CreateUser = true

	MainMenuIndex(w, data)

	data.CreateUser = false
	//tmpl, _ := template.ParseFiles("index/signIn.html")

	//tmpl.Execute(w, data)
}

// Submit the creation of user
func createUserSubmitHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	username := r.FormValue("username")
	var createUser bool
	if username != "" || password != "" {
		createUser = jsonMod.SaveUserJson(username, password)
	}

	if createUser {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/create_user", http.StatusSeeOther)
	}

}

func authHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	username := r.FormValue("username")

	authSuccess := userAuth.AuthenticateUser(username, password)

	if authSuccess {
		fmt.Println("login succeeded")

		UserData, _ := jsonMod.ReadUserJson(jsonMod.JsonPathUser)

		uuID := UserData.User[username]["uuid"]

		userID := 123

		session, _ := store.Get(r, "session-name")
		session.Values["user_id"] = userID
		session.Values["uuID"] = uuID
		session.Values["username"] = username
		session.Save(r, w)

		// Generate a session token
		sessionID := generateSessionID()

		// Store the session token in a cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "sessionID",
			Value:   sessionID,
			Expires: time.Now().Add(24 * time.Hour), // Set cookie expiration time
			Path:    "/",
		})

		http.Redirect(w, r, "/", http.StatusFound)

		data.FailedLoginAttempt = false

		data.CreateUser = false
	} else {
		fmt.Println("login failed")
		//http.Redirect(w, r, "/", http.StatusFound)
		data.FailedLoginAttempt = true

		data.CreateUser = false

		MainMenuIndex(w, data)

	}

}

func main() {
	//db := database.Connect()
	//database.CreateWordListTable(db)
	//id := database.InsertNewWordList(db)
	//fmt.Println(id)
	//database.ReadWord(db)

	http.Handle("/theme/", http.StripPrefix("/theme/", http.FileServer(http.Dir("theme"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))

	http.HandleFunc("/", handler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/create_user", createUserHandler)
	http.HandleFunc("/create_user_submit", createUserSubmitHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/question", questionHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/next", nextHandler)
	http.HandleFunc("/prev", prevHandler)
	http.HandleFunc("/wordlist_changed", wordlistChangeHandler)
	http.HandleFunc("/wordcount_changed", wordcountChangeHandler)
	http.HandleFunc("/language_changed", languageChangeHandler)
	http.HandleFunc("/word_result", wordResultHandler)
	http.HandleFunc("/add_word", addWordMainHandler)
	http.HandleFunc("/add_word_save", addWordListHandler)
	http.HandleFunc("/add_word_final", addWordsHandler)
	http.ListenAndServe(":8080", nil)

}
