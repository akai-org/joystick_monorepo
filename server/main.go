package main

import (
	"akai.org.pl/joystick_server/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ctrl := controller.New()
	prepareSocket()
	fmt.Println("Began listening...")
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/join", ctrl.RegisterNewPlayer)
	http.HandleFunc("/create", ctrl.CreateRoom)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
