package main

import (
	"azla_go_learning/internal/json"
	"azla_go_learning/internal/userAuth"
	"azla_go_learning/internal/viewData"
	"azla_go_learning/internal/words"
	"azla_go_learning/internal/routes"
	//"azla_go_learning/internal/database"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// when you first start the quiz
func questionHandler(w http.ResponseWriter, r *http.Request) {
	viewData.Data.CurrentQuestion = 0
	viewData.Data.MaxAmountOfWords = 0
	viewData.Data.CorrectAnswers = 0
	viewData.Data.InCorrectAnswers = 0
	viewData.Data.IsComplete = map[string]bool{}
	viewData.Data.CorrectAnswersList = map[string]string{}
	viewData.Data.InCorrectAnswersList = map[string]string{}
	viewData.Data.AvailableWords = []string{}

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

	viewData.Data.WordList = wordlist

	var words = []string{}
	var correct = []string{}

	if r.FormValue("examMode") != "" {
		viewData.Data.ExamMode = true
		viewData.Data.ExamModeAction = "/next"
		viewData.Data.ExamModeString = "Next"
	} else {
		viewData.Data.ExamMode = false
		viewData.Data.ExamModeAction = "/submit"
		viewData.Data.ExamModeString = "Submit"
	}

	selectedWordList := r.FormValue("wordListOpt")
	selectedLanguage := r.FormValue("languageOpt")
	selectedCount := r.FormValue("wordCountOptions")

	viewData.Data.SelectedWordList = selectedWordList
	viewData.Data.SelectedLanguage = selectedLanguage

	// Convert selected word count to int
	viewData.Data.MaxAmountOfWords = convertToNum(selectedCount)

	// Create the word and correct slices/arrays
	for key, value := range viewData.Data.WordList[selectedWordList] {
		switch viewData.Data.SelectedLanguage {
		case "Azerbajani":
			words = append(words, key)
			correct = append(correct, value)
		case "English":
			words = append(words, value)
			correct = append(correct, key)
		}
	}

	// Current word
	currentWord := words[viewData.Data.CurrentQuestion]
	currentIndex := viewData.Data.CurrentQuestion + 1
	
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

	viewData.Data.Words = words
	viewData.Data.Correct = correct
	viewData.Data.WordImage = selectedUrl

	if viewData.Data.MaxAmountOfWords >= len(viewData.Data.Words) {
		fmt.Println("Wordlist doesn't have enough words:", len(viewData.Data.Words))
		fmt.Println("Choosen option:", viewData.Data.MaxAmountOfWords)
	}

	for i := 0; i < len(viewData.Data.Words); i++ {
		if i > viewData.Data.MaxAmountOfWords {
			break
		} else {
			viewData.Data.AvailableWords = append(viewData.Data.AvailableWords, viewData.Data.Words[i])
		}
	}

	fmt.Println(viewData.Data.AvailableWords)

	if viewData.Data.ExamMode {

		viewData.Data.CurrentWord = currentWord
		viewData.Data.CurrentIndex = currentIndex
		tmpl, _ := viewData.CreateQuestionTemp()

		//tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, viewData.Data)

	} else {

		//htmlStr := CreateQuestionTemp(currentIndex, currentWord, data, "/submit")
		tmpl, err := viewData.CreateQuestionTemp()

		//tmpl, _ := template.New("t").Parse(htmlStr)
		viewData.Data.CurrentWord = currentWord
		viewData.Data.CurrentIndex = currentIndex

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, viewData.Data)

	}

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

	for key, value := range viewData.Data.InCorrectAnswersList {
		fmt.Println(key, value)
	}

	tmpl, _ := template.ParseFiles("index/resultMenu.html")

	tmpl.Execute(w, viewData.Data)
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


// Launch the login screen
func loginHandler(w http.ResponseWriter, r *http.Request) {
	viewData.Data.CreateUser = false

	http.Redirect(w, r, "/", http.StatusFound)
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
	viewData.Data.CreateUser = true

	routes.MainMenuIndex(w, viewData.Data)

	viewData.Data.CreateUser = false
	//tmpl, _ := template.ParseFiles("index/signIn.html")

	//tmpl.Execute(w, viewData.Data)
}

// Submit the creation of user
func createUserSubmitHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password-confirm")
	username := r.FormValue("username")
	viewData.Data.CreateUserMes = ""
	var createUser bool

	if username == "" || password == "" {
		viewData.Data.CreateUserMes = "Username or Password can't be empty"
	}
	if password != passwordConfirm {
		viewData.Data.CreateUserMes = "Passwords Doesn't Match"

	} else if username != "" || password != "" {
		createUser = jsonMod.SaveUserJson(username, password)
	}

	if createUser {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if viewData.Data.CreateUserMes == "" {
		viewData.Data.CreateUserMes = "Username is taken"
		http.Redirect(w, r, "/create_user", http.StatusSeeOther)
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

		viewData.Data.FailedLoginAttempt = false

		viewData.Data.CreateUser = false
	} else {
		fmt.Println("login failed")
		//http.Redirect(w, r, "/", http.StatusFound)
		viewData.Data.FailedLoginAttempt = true

		viewData.Data.CreateUser = false

		routes.MainMenuIndex(w, viewData.Data)

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

	http.HandleFunc("/", routes.MainHandler)
	http.HandleFunc("/api/words", routes.WordsApiHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/create_user", createUserHandler)
	http.HandleFunc("/create_user_submit", createUserSubmitHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/question", questionHandler)
	http.HandleFunc("/submit", routes.SubmitHandler)
	http.HandleFunc("/next", routes.NextHandler)
	http.HandleFunc("/prev", routes.PrevHandler)
	http.HandleFunc("/wordlist_changed", wordlistChangeHandler)
	http.HandleFunc("/wordcount_changed", wordcountChangeHandler)
	http.HandleFunc("/language_changed", languageChangeHandler)
	http.HandleFunc("/word_result", wordResultHandler)
	http.HandleFunc("/add_word", addWordMainHandler)
	http.HandleFunc("/add_word_save", addWordListHandler)
	http.HandleFunc("/add_word_final", addWordsHandler)
	http.ListenAndServe(":8080", nil)

}
