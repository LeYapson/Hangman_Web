package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ouais, Champion, c, comment?")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tu, veux, jouer?")
}


func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("Play", Play)

	fmt.Println("http://localhost:8080 - Server started on port", port)
http.ListenAndServe(port, nil)
}
