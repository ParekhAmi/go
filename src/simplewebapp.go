package main

import (
	"log"
	"net/http"
)

/*func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang basic web app learning..")
}*/

func home(w http.ResponseWriter, r *http.Request) {

	// Check if the current request URL path exactly matches "/".
	//If it doesn't, use the http.NotFound() function to send a 404 response to the client.

	// Like.... http://localhost:8000/missing. You should get a 404 page not found

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
	/* http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil) */

}
