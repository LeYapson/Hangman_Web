package main

import (
	"fmt"
	"net/http"
)

func localhost() {
	fmt.Print("starting a local host on port 8080.", "\n")
	fmt.Print("\n", "enjoy your game 0w0", "\n")
	http.ListenAndServe(":8080", nil)
}
