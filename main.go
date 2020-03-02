package main

import (
	"log"
	"net/http"
	"os"

	"./db"
)

type Env struct {
	db db.DataStore
}

var mySigningKey = []byte(os.Getenv("AUTH_SERVER_SECRET"))

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	handleRequests()
}
