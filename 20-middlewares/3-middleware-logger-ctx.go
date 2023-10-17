package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid" //external package to get this package use: go get github.com/google/uuid
)

// reqKey is a type used for context keys.
type reqKey int

// RequestIDKey is a key for 'request id' context value.
const RequestIDKey reqKey = 123

func main() {
	// Setup the /home route, applying middleware to the handler.
	http.HandleFunc("/home", RequestIdMid(LoggingMid(Midware(homePage))))
	// Start listening and serving requests on port 8080.
	http.ListenAndServe(":8081", nil)
}

// homePage is a simple HTTP handler function that writes a message to the response.
func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("In home Page handler")
	fmt.Fprintln(w, "this is my home page.")
}

// RequestIdMid adds a request ID to the context of every request.
func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate a new unique ID.
		uuid := uuid.NewString()
		// Create a new request context with the ID.
		ctx := context.WithValue(r.Context(), RequestIDKey, uuid)
		// Pass the responsibility to the next Middleware/HandlerFunc in the chain with the new context.
		next(w, r.WithContext(ctx))
	}
}

// LoggingMid is a HTTP Middleware that logs the start and end of each request along with the request ID.
func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the request ID from the context.
		reqId, ok := r.Context().Value(RequestIDKey).(string)
		if !ok {
			reqId = "Unknown"
		}
		// Print log with request ID, method and path.
		log.Printf("%s : started   : %s %s ", reqId, r.Method, r.URL.Path)
		// Ensure log of completion when the request is done.
		defer log.Println("completed")
		// Pass the responsibility to the next HandlerFunc in the chain.
		next(w, r)
	}
}

func Midware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId, ok := r.Context().Value(RequestIDKey).(string)
		if !ok {
			reqId = "Unknown"
		}
		log.Printf("Request id from context= %s", reqId)
		defer log.Println("Mid1 Complete")

		next(w, r)

	}
}
