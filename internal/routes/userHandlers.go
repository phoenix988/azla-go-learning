package routes

import (
	"net/http"
	"fmt"
	"time"
	"azla_go_learning/internal/userAuth"
	"azla_go_learning/internal/json"
	"azla_go_learning/internal/viewData"
	"azla_go_learning/internal/char"

)


// Launch the login screen
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	viewData.Data.CreateUser = false

	http.Redirect(w, r, "/", http.StatusFound)
}


// Handle logout event
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := viewData.Store.Get(r, "session-name")

	// Revoke authentication
	delete(session.Values, "user_id")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}


// Create the user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	viewData.Data.CreateUser = true

	MainMenuIndex(w, viewData.Data)

	viewData.Data.CreateUser = false
	//tmpl, _ := template.ParseFiles(viewData.TemplatePath+"signIn.html")

	//tmpl.Execute(w, viewData.Data)
}

// Submit the creation of user
func CreateUserSubmitHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password-confirm")
	username := r.FormValue("username")
	viewData.Data.CreateUserMes = ""
	var createUser bool

	if username == "" {
		viewData.Data.CreateUserMes = "Username or Password can't be empty"
	} else if password == "" {
		viewData.Data.CreateUserMes = "Username or Password can't be empty"
	} else if password != passwordConfirm {
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



func AuthHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	username := r.FormValue("username")

	authSuccess := userAuth.AuthenticateUser(username, password)

	if authSuccess {
		fmt.Println("login succeeded")

		UserData, _ := jsonMod.ReadUserJson(jsonMod.JsonPathUser)

		uuID := UserData.User[username]["uuid"]

		userID := 123

		session, _ := viewData.Store.Get(r, "session-name")
		session.Values["user_id"] = userID
		session.Values["uuID"] = uuID
		session.Values["username"] = username
		session.Save(r, w)

		// Generate a session token
		sessionID := char.GenerateSessionID()

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

		MainMenuIndex(w, viewData.Data)

	}

}


func NewUserRoutes() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/create_user", CreateUserHandler)
	http.HandleFunc("/create_user_submit", CreateUserSubmitHandler)
	http.HandleFunc("/auth", AuthHandler)

}
