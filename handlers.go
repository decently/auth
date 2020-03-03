package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (env *Env) createAccount(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" Wrong header", http.StatusBadRequest)
		return
	}

	acc := Account{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" Bad Io Util conversion", http.StatusInternalServerError)
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
func (env *Env) requestToken(w http.ResponseWriter, r *http.Request) {

}
func (env *Env) requestNewApp(w http.ResponseWriter, r *http.Request) {

}
