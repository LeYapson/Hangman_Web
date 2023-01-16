package main

import (
	//"github.com/sinjin314/hangman-classic"
	"net/http"
)

func main() { // permet de lancer le localhost et charger les pages web
	http.HandleFunc("/", Home)
	http.HandleFunc("/game", Game)
	localhost()

}
