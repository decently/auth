package main

import (
	"log"
	"net/http"
	"os"

	"github.com/djbrunelle/auth/service"
	"github.com/gorilla/mux"
)

type Env struct {
	db service.DataStore
}

var (
	mySigningKey = []byte(os.Getenv("APP_KEY"))
)

func handleRequests() {
	db, err := service.OpenDB(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	env := &Env{db}

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.Path("/accounts").HandlerFunc(env.createAccount)
	subRouter.Path("/token").HandlerFunc(env.requestToken)
	log.Fatal(http.ListenAndServe(":9090", router))
}

func main() {
	handleRequests()
}
