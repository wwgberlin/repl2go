package main

import (
	"github.com/wwgberlin/repl2go/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/run", handler.RunHandler)
	fmt.Println("Starting server")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
