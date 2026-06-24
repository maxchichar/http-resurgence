// r.Header.Get() returns an empty string for a header key that was never sent.

package main

import(
	"fmt"
	"log"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request)  {
	header := r.Header.Get("X-Custom-Token")
	if header == ""{
		w.WriteHeader(http.StatusBadRequest) // Return code 404 bad request
		w.Write([]byte("X-Custom-Token header is missing")) // Error fallback message if header is empty
		return
	}

	content := r.Header.Get("Content-Type")
	var types string
	switch content {
	case "":
		types = "Content-Type not provided" // if content empty return this message
	default:
		types = fmt.Sprintf("Content-Type: %s", content) // building string
	}
	

	token := fmt.Sprintf("Token received: %s", header) // building string

	fmt.Fprintf(w, "%s\n%s", token, types) // Printing the output to the client
}

// Hosting server at :8080
func main()  {
	http.HandleFunc("/headers", headerHandler)
	fmt.Println("Server running at http://localost:8080/header")
	log.Fatal(http.ListenAndServe(":8080", nil))
}