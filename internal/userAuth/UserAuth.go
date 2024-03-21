package userAuth

import (
	//"golang.org/x/crypto/bcrypt"
	"azla_go_learning/internal/json"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(username, password string) bool {
	type Auth struct {
		CheckUser bool
	}

	authData := Auth{}
	authData.CheckUser = false

	userData, _ := jsonMod.ReadUserJson(jsonMod.JsonPathUser)

	for key := range userData.User {
		if key == username {
			authData.CheckUser = true
		}
	}

	hashedPassword := userData.User[username]["password"]

	if authData.CheckUser == true {
		// Compare the hashed password with the provided password
		comparePassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if comparePassword == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

	return false

}
