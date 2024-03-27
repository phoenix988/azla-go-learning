package main

import (
	"azla_go_learning/internal/routes"
	"azla_go_learning/internal/cmd"
	"net/http"
	"log"
	"fmt"
)



func main() {
	// Port to use
	var AzlaConfig = cmd.ParseFlag()

	// Read Style.css and script.js
	http.Handle("/theme/", http.StripPrefix("/theme/", http.FileServer(http.Dir("theme"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))

	routes.NewMainRoutes() // Main routes for the main quiz
	routes.NewWordRoutes() // wordlist related routes
	routes.NewUserRoutes() // User related routes
	routes.NewUpdateStringRoutes() // Update strings dynamically
	
	// Api listener
	http.HandleFunc("/api/words", routes.WordsApiHandler) // Word api
	fmt.Println("Listening on Port: " +AzlaConfig.Port)
	log.Fatal(http.ListenAndServe(":"+AzlaConfig.Port, nil))

}
