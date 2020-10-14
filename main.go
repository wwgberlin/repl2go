package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Hist History
	Last string // "\n or \r\n " separated or whatever browser do to inputs
}

type History []struct {
	Inputs string
	Output []string //todo - this depends on what we get from the interpreter
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the history from the message
	// Instaniate Interpreter instance : https://github.com/traefik/yaegi/tree/e32da38ad0fa87aced9812c391b57853225cb344#as-an-embedded-interpreter
	// Iterate over the history
	// We would extract the below to a new packages that wraps yaegi and also handles classifying commands (ie import statements))
	// Invoke Interpreter(Yaegi) to evaluate each command in the history
	// (BONUS) Compare return values from Yaegi with historical output
	// Invoke Interpreter(Yaegi) on the Message.Last command
	// Store the latest output
	// Respond 200 where the body contains a new json with all the historical outputs and (at least) latest output & a dummy string for the valid go file
}

func main() {
	http.HandleFunc("/run", runHandler)
	fmt.Println("Starting server")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
