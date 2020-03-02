package main

import (
	"fmt"
	"net/http"
)

func generateHash(password string) string {

}

func token(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")

}
