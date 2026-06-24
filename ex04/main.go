package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost{ // Rejecting non-POST request with 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Check for errors immediately
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userName := r.FormValue("username")
	if userName == "" { // Checking for empty username
		w.WriteHeader(http.StatusBadRequest) // returns a 400 code as Bad request
		w.Write([]byte("username is required")) // message sent to client
		return
	}

	lang := r.FormValue("language")
	if lang == "" { // Checking for empty language
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("language is required"))
		return
	}
	
	// Prints a greeting message to w with the given userName and language.
	fmt.Fprintf(w, "Hello %s, you are coding in %s!", userName, lang)
}

func main()  {
	http.HandleFunc("/form", formHandler)
	fmt.Println("Server running at http://localhost:8080/form")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
