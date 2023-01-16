package main

import (
	//"github.com/sinjin314/hangman-classic"
	"fmt"
	"net/http"
	"text/template"
)

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/game", Game)
	localhost()
	
}
func Home (w http.ResponseWriter, r *http.Request){
	Rendertemplate(w, "home")
}

func Game (w http.ResponseWriter, r *http.Request) {
	Rendertemplate(w, "game")
}

func Rendertemplate(w http.ResponseWriter, tmpl string) {
	t, err :=template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		fmt.Println("error")
	}
	t.Execute(w, nil)
}