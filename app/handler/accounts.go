package handler

import (
	"encoding/json"
	"github.com/carlosdamazio/Stone-REST-API/app/model"
	"github.com/carlosdamazio/Stone-REST-API/app/serializer"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetAccounts(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	accounts := model.GetAccounts(db)

	if accounts == nil {
		respondJsonError(w, http.StatusNotFound,"No account was found in the database.")
		return
	} else {
		respondJson(w, http.StatusOK, accounts)
	}
}

func GetAccountBalance(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	account := &model.Account{}
	vars := mux.Vars(r)
	id := vars["account_id"]

	err := account.GetAccount(db, id)

	if err != nil {
		respondJsonError(w, http.StatusNotFound, err.Error())
		return
	} else {
		respondJson(w, http.StatusOK, map[string]float64{
			"balance": account.Balance,
		})
	}
}

func CreateAccount(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	account := &model.Account{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(account); err != nil {
		respondJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok, err := serializer.SerializeAccountRequest(account); !ok {
		respondJsonError(w, http.StatusBadRequest, err)
		return
	}

	account.Save()

	res, err := account.InsertAccount(db)

	if err != nil {
		respondJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusCreated, res)
}
