package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/djbrunelle/auth/service"
)

func (env *Env) createAccount(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" Wrong header", http.StatusBadRequest)
		return
	}

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	acc := service.Account{}

	acc.Email = pair[0]

	hash, err := generateHash(pair[1])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" hash issue", http.StatusInternalServerError)
		return
	}

	acc.HashedPassword = string(hash)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" Bad IO Util conversion", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(body), &acc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" Bad Json Unmarshal", http.StatusBadRequest)
		return
	}

	err = env.db.CreateAccount(acc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" "+err.Error(), http.StatusBadRequest)
		return
	}
}

func (env *Env) updateAccount(w http.ResponseWriter, r *http.Request) {

}

func (env *Env) deleteAccount(w http.ResponseWriter, r *http.Request) {

}

func (env *Env) requestToken(w http.ResponseWriter, r *http.Request) {

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	acc, err := env.db.GetAccount(pair[0])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" "+err.Error(), http.StatusBadRequest)
		return
	}

	err = validateHash(pair[1], acc.HashedPassword)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized)+" "+err.Error(), http.StatusUnauthorized)
		return
	}

	token, _ := generateToken(pair[0])

	w.Write([]byte(token))

}

func (env *Env) authenticate(w http.ResponseWriter, r *http.Request) {

}
func (env *Env) requestNewApp(w http.ResponseWriter, r *http.Request) {

}
