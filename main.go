package main

import (
	//"github.com/sinjin314/hangman-classic"
	"net/http"
)

func main() { // permet de lancer le local host et charger les page web
	http.HandleFunc("/", Home)
	http.HandleFunc("/game", Game)
	localhost()

}
