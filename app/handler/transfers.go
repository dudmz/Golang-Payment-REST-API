package handler

import (
	"encoding/json"
	"github.com/carlosdamazio/Stone-REST-API/app/model"
	"github.com/carlosdamazio/Stone-REST-API/app/serializer"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetTransfers(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	transfers := model.GetTransfers(db)

	if transfers == nil {
		respondJsonError(w, http.StatusNotFound,"No transfer was found in the database.")
		return
	} else {
		respondJson(w, http.StatusOK, transfers)
	}
}

func MakeTransfer(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	originAccount := model.Account{}
	destinationAccount := model.Account{}
	transfer := &model.Transfer{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&transfer); err != nil {
		respondJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok, err := serializer.SerializeTransferRequest(transfer); !ok {
		respondJsonError(w, http.StatusBadRequest, err)
		return
	}

	transfer.Save()

	errOrigin := originAccount.GetAccount(db, transfer.OriginAccount)
	errDestination := destinationAccount.GetAccount(db, transfer.DestinationAccount)

	if errOrigin != nil {
		respondJsonError(w, http.StatusBadRequest, "Origin account doesn't exist.")
		return
	}

	if errDestination != nil {
		respondJsonError(w, http.StatusBadRequest, "Destination account doesn't exist.")
		return
	}

	res, err := transfer.MakeTransfer(db, &originAccount, &destinationAccount)

	if err != "" {
		respondJsonError(w, http.StatusBadRequest, err)
		return
	}

	respondJson(w, http.StatusCreated, res)
}
