package database


import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"encoding/json"
	"fmt"

)


type wordsData struct {
	Wordlist map[string]map[string]map[string]string `json:"wordlist"`
	Username string
}

// Connect to the database
func Connect() (*sql.DB) {
	connStr := "postgres://postgres:password@192.168.0.4:5432/azla_go_learning?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db

}

func Close(db *sql.DB) {
	db.Close()
}

func CreateWordListTable(db *sql.DB) {
	
	query := `CREATE TABLE IF NOT EXISTS word_list (
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	wordlist_data JSONB
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

}


// Insert the new words to the database
func InsertNewWordList(db *sql.DB) int {
	wordlist := wordsData {
		Username: "user2",
		Wordlist: map[string]map[string]map[string]string{},
	}

	wordlist.Wordlist["user1"] = map[string]map[string]string{}

	wordlist.Wordlist["user1"]["list1"] = map[string]string{
		"word1": "definition1",
	}

	jsonData, _ := json.Marshal(wordlist)
	
	query := `INSERT INTO word_list (username, wordlist_data)
			VALUES ($1, $2) RETURNING id`

	var pk int
	
	err := db.QueryRow(query, wordlist.Username, string(jsonData)).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}

	return pk
}

func ReadWord(db *sql.DB) {
	query := `SELECT * FROM word_list WHERE username = $1`

	rows, _ := db.Query(query, "user1")


	  // Iterate over the rows
    for rows.Next() {
        var id int
        var storedUsername string
        var storedData string

        // Print the data if the username matches
        if storedUsername == "user1" {
            fmt.Printf("ID: %d, Username: %s, Data: %s\n", id, storedUsername, storedData)
        }
    }


}
