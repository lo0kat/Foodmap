package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint hit : HomePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/checkpoints", returnAllCheckpoints)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Checkpoint struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Address     string   `json:"address"`
	Description string   `json:"desc"`
	Rating      uint     `json:"rating"`
}

var Checkpoints []Checkpoint

func main() {
	Checkpoints =
		append(Checkpoints, Checkpoint{Name: "EGGDROP", Tags: []string{"Restaurant", "Ham cheese Sandwich"}, Address: "AjouUniv", Description: "Good place to eat", Rating: 5}, Checkpoint{Name: "MRCHEF", Tags: []string{"Restaurant", "Pork"}, Address: "AjouUniv", Description: "Good place to eat pork", Rating: 4})
	handleRequests()
}

func returnAllCheckpoints(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point Hit : returnAllCheckpoints")
	json.NewEncoder(w).Encode(Checkpoints)
}
