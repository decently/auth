package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func generateHash(password string) []byte {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	fmt.Println(err)

	return hashedPassword
}

func token(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")

}
