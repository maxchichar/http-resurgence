package main

import(
	"io"
	"fmt"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST"{ // Rejecting any non-POST request
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body) // Reading the full body request 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // This run's when the function  returns, it's like a sticky note not to forget what to do.

	if len(body) == 0{ // Checking length of body
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body cannot be empty")) // Error message
		return
	}

	fmt.Fprintf(w, "%s", body) // Writing the body content back to w exactly as recieved.
}

// Server Setup
func main()  {
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("Server running at http://localhost:8080/echo")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
