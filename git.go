package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := "index.html"
	if r.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(r.URL.Path)
	}

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func main() {
	fmt.Println("http server up at port 8081!")
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8081", nil)
}
