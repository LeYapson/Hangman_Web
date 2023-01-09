package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("starting a local host on port 8080.","\n")
	fmt.Print("\n","enjoy your game 0w0")
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}