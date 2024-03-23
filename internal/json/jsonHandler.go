package jsonMod

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var JsonPath = "data/data.json"
var JsonPathUser = "data/users.json"


type wordsData struct {
	Wordlist map[string]map[string]map[string]string `json:"wordlist"`
}

type User struct {
	User map[string]map[string]string
}

// Function to read words from json data file
func ReadWordJson(jsonPath string, username string) (wordsData, error) {
	// Decode JSON data into a Go data structure
	var data wordsData

	// Open the JSON file for reading
	file, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return data, err
	}

	return data, err
}

// Function to read words from json data file
func ReadUserJson(jsonPath string) (User, error) {
	// Decode JSON data into a Go data structure
	var data User

	// Open the JSON file for reading
	file, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return data, err
	}

	return data, err
}

// Function to save words to json datta file
func SaveWordJson(wordListName string, english []string, aze []string, username string) {
	// Sample data
	wordData := wordsData{
		Wordlist: map[string]map[string]map[string]string{},
	}

	wordData.Wordlist[username] = map[string]map[string]string{}
	wordData.Wordlist[username][wordListName] = map[string]string{}

	for i := 0; i < len(english) && i < len(aze); i++ {
		// Add to the map
		wordData.Wordlist[username][wordListName][english[i]] = aze[i]
	}

	// Read existing data from file
	existingData, existingErr := ReadWordJson(JsonPath, username)

	if existingErr != nil {
		fmt.Println("Error reading existing data:", existingErr)
	} else {

		for key, value := range wordData.Wordlist {
			
			// Merge existing data with new data
			existingData.Wordlist[key][wordListName] = value[wordListName]
		}

	}

	// Marshal data into JSON
	var jsonData []byte
	var err error
	if existingErr != nil {
		jsonData, err = json.MarshalIndent(wordData, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

	} else {
		jsonData, err = json.MarshalIndent(existingData, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
	}

	// Open a file for writing (create it if it doesn't exist)
	file, err := os.Create(JsonPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Printf("Data written to %s successfully!", JsonPath)

}

// Function to save words to json datta file
func SaveUserJson(username string, password string) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Save the data
	userData := User{}

	userData.User = map[string]map[string]string{}

	userData.User[username] = map[string]string{
		"username": username,
		"password": string(hashedPassword),
		"uuid":     uuid.NewString(),
	}

	// Read existing data from file
	existingData, existingErr := ReadUserJson(JsonPathUser)

	if existingErr != nil {
		fmt.Println("Error reading existing data:", existingErr)
	} else {

		for key := range existingData.User {

			if username == key {
				fmt.Println("User Exist")
				return false
			}

		}

		// Merge existing data with new data
		for key, value := range userData.User {
			existingData.User[key] = value
		}

	}

	// Marshal data into JSON
	var jsonData []byte
	var err error
	if existingErr != nil {
		jsonData, err = json.MarshalIndent(userData, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return false
		}

	} else {
		jsonData, err = json.MarshalIndent(existingData, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return false
		}
	}

	// Open a file for writing (create it if it doesn't exist)
	file, err := os.Create(JsonPathUser)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer file.Close()

	// Write JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return false
	}

	fmt.Printf("Data written to %s successfully!", JsonPathUser)

	return true

}
