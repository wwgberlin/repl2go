package main

import (
	"github.com/wwgberlin/repl2go/handler"
	"log"
	"net/http"
	"time"
)

// logging middleware that logs the incoming requests
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// log request
		log.Printf("%s %s on %s%s\n", r.Proto, r.Method, r.Host, r.URL)

		// hand over to next handler
		next.ServeHTTP(w, r)

		// log response
		end := time.Now()
		duration := end.Sub(start)
		log.Printf("done in %fs!", duration.Seconds())
	})
}

func main() {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(handler.RunHandler)
	mux.Handle("/run", logger(finalHandler))
	mux.Handle("/", logger(http.NotFoundHandler()))

	log.Println("Starting server on localhost:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
