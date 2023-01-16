package main

import (
	//"github.com/sinjin314/hangman-classic"
	"fmt"
	"net/http"
	"text/template"
)

func main(){ // permet de lancer le local host et charger les page web
	http.HandleFunc("/", Home)
	http.HandleFunc("/game", Game)
	localhost()
	
}
func Home (w http.ResponseWriter, r *http.Request){ // afiche les page
	Rendertemplate(w, "home")
}

func Game (w http.ResponseWriter, r *http.Request) { // Hangman web
	Rendertemplate(w, "game")
}

func Rendertemplate(w http.ResponseWriter, tmpl string) { // Cherche le rendu de la template dans les fichier
	t, err :=template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		fmt.Println("error")
	}
	t.Execute(w, nil)
}