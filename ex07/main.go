// template.Must panics if the template string is broken

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
	<h1>{{.Title}}</h1>
	<p>{{if eq .Style "bold"}}<strong>{{.Body}}</strong>{{else}}{{.Body}}{{end}}</p>
</body>
</html>

`

type PageData struct{
	Title string
	Body string
	Style string
}

var tmpl = template.Must(template.New("page").Parse(tmplStr))

func renderHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Query().Get("title")
	if title == ""{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("title and body are required"))
		return
	}

	body := r.URL.Query().Get("body")
	if body == ""{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("title and body are required"))
		return
	}

	style := r.URL.Query().Get("style")

	w.Header().Set("Content-Type", "text/html")

	err := tmpl.Execute(w, PageData{Title: title, Body: body, Style: style})
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("template execution failed"))
		return
	}

}

func main()  {
	http.HandleFunc("/render", renderHandler)
	fmt.Println("Server running at http://localhost:8080/render")
	log.Fatal(http.ListenAndServe(":8080", nil))
}