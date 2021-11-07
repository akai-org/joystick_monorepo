package main

import (
	"akai.org.pl/joystick_server/communication"
	"fmt"
	"log"
	"net/http"
)

var counter uint8 = 0

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func main() {
	prepareSocket()
	fmt.Println("Began listening...")
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	http.HandleFunc("/join", communication.RegisterNewPlayer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}