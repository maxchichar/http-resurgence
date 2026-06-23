package main 

import(
	"fmt"
	"log"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request)  {
	// Fprintf writes to w a message indicating the request method
	fmt.Fprintf(w, "You made a %s request", r.Method)
}

func main()  {
	http.HandleFunc("/method-inspector", methodHandler) // Registers methodHandler to handle HTTP requests at "/method-inspector".
	fmt.Println("Server running at http://localhost:8080/method-inspector")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

