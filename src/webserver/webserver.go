package main

import (
	"log"
	"net/http"
)

func runServer() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on localhost:3000...")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runServer()
}
