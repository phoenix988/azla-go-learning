package main

// Static variables I set here
// Plus the main pagedata I want to pass to the template

import (
	"azla_go_learning/internal/json"
	//"azla_go_learning/internal/viewData"
	"github.com/gorilla/sessions"
)

// Language options
var languageOptions []string = []string{"Azerbajani", "English"}

// Set the wordlist count options
var amountOfWords = []int{5, 10, 15, 20, 25, 30} // Amount of words to choose from

// JsonPath
var jsonPath = jsonMod.JsonPathUser
var jsonPathUser = jsonMod.JsonPath
var templatePath = "templates/"

// Create session store
var store = sessions.NewCookieStore([]byte("secret-key"))

// User data struct
type User struct {
	ID       int
	Username string
	Password string // Hashed password
}

