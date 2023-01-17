package main

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) { // affiche la page d'accueil
	Rendertemplate(w, "home")
}

func Game(w http.ResponseWriter, r *http.Request) { // affiche la page contenant le jeu
	Rendertemplate(w, "game")
}

func Rendertemplate(w http.ResponseWriter, tmpl string) { // Cherche le rendu des templates dans les fichiers
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
