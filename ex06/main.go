// If i use mainMux.Handle("/api", ...) without the trailing slash, i have zero access to those page which in turn returns a 404 page not found.

// Can't work cause there would be error and also if i'm to answer that i would say nothing only the /api.

// So Go see's mainMux, Checks "does it have *http.ServeHTTP?" and answers yes and accept's it.

package main

import(
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "pong")
}

func greetHandler(w http.ResponseWriter, r *http.Request)  {
	var name string
	name = r.URL.Query().Get("name")
	if name == ""{
		name = "Stranger"
	}
	fmt.Fprintf(w, "Greetings, %s!", name)
}

func main()  {
	// Creating a new mux instead of using the default one
	apiMux := http.NewServeMux()

	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	mainMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mainMux))
}