package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetTransfers(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, map[string]string{
		"msg": "All transfers will appear in here!",
	})
}

func MakeTransfer(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, map[string]string{
		"msg": "You'll be able to make a transfer in here!",
	})
}
