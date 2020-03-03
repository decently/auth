package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Env struct {
	db DataStore
}

var mySigningKey = []byte(os.Getenv("AUTH_SERVER_SECRET"))

func handleRequests() {
	db, err := OpenDB(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	env := &Env{db}

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.Path("/accounts").HandlerFunc(env.createAccount)
	log.Fatal(http.ListenAndServe(":9090", router))
}

func main() {
	handleRequests()
}
