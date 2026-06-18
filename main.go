package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice contaning
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// and try to convert it to an interger using the srtconv.Atoi() function. If
	// it can´t be converted to an interger, or the value is less than 1, we
	// return a 404 page not found response.

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmft.Sprintf() function to interpolate the id value with a
	// message, then write it as the HTTP response.
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Add a snippetCrate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating anew Snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to inittialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home) // Now restrict this route to exact matches on / only.
	// Add the two news handlers
	mux.HandleFunc("/snippet/view/{id}", snippetView) // Now add the {id} wildcard segment
	mux.HandleFunc("/snippet/create", snippetCreate)
	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created.If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and terminate the
	// program. Note tha any error returned by http.ListenAndServe() is always
	// non-nil.

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
