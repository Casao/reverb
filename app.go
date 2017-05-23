package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/satori/go.uuid"
)

type ApiResponse struct {
	ID      string
	Headers http.Header
	Body    string
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.NewV4().String()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	headers := r.Header
	ar := ApiResponse{ID: uuid, Body: body, Headers: headers}
	JSONResponse(w, ar, 200)
}

func main() {
	http.HandleFunc("/", handlePost)
	http.ListenAndServe(":8080", nil)
}

func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		Logger.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

// Logger is used to send logging messages to stdout.
var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)
