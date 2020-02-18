package handler

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetAccounts(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, map[string]string{
		"msg": "All accounts will appear in here!",
	})
}

func GetAccountBalance(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["account_id"]

	respondJson(w, http.StatusOK, map[string]string{
		"msg": "Balance of account " + id + " will appear in here!",
	})
}

func CreateAccount(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, map[string]string{
		"msg": "You'll create an account in here!",
	})
}
