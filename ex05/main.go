// When i swapped w.WriteHeader with fmt.Fprintf in the last block of the code, i observed that the status code in the body and the actuall HTTP status are completely out of sync.

package main

import(
	"fmt"
	"log"
	"strconv"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request)  {
	code := r.URL.Query().Get("code") // Reading the code query parameter
	// If code is empty return error 400
	if code == ""{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("code parameter is required"))
		return
	}

	statusCode, err := strconv.Atoi(code) // Conversion of code to integer
	// Error handling for invalid code
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("code must be a valid integer"))
		return
	}

	if statusCode < 100 || statusCode > 599 { // Code must be between 100 and 599
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("code must be a valid HTTP status code (100–599)"))
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Responding with status %d %s", statusCode, http.StatusText(statusCode))
}

// Server Hosting at :8080
func main()  {
	http.HandleFunc("/status", statusHandler)
	fmt.Println("Server running at http://localhost:8080/status")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
