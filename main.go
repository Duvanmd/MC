package main

import (
	"adn/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	//EndPoind

	mux.HandleFunc("/api/results", handlers.GetAdn).Methods("GET")
	mux.HandleFunc("/api/mutant", handlers.CreateAdn).Methods("POST")

	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
