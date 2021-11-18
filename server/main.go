package main

import (
	"akai.org.pl/joystick_server/restapi"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func main() {
	prepareSocket()
	fmt.Println("Began listening...")
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	http.HandleFunc("/join", restapi.RegisterNewPlayer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
