package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getMessage(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, os.Getenv("Message")+"\n")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", getMessage)
	err := http.ListenAndServe(":"+os.Getenv("Port"), nil)
	if err != nil {
		log.Println(err)
	}
}
