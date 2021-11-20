package main

import (
	"akai.org.pl/joystick_server/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	prepareSocket()
	fmt.Println("Began listening...")
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/join", controller.RegisterNewPlayer)
	http.HandleFunc("/create", controller.CreateRoom)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
