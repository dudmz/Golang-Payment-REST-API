package app

import (
	"context"
	"github.com/carlosdamazio/Stone-REST-API/app/handler"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router *mux.Router
	DB	   *mongo.Client
}

func (a *App) Initialize() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("Application could not connect to database.\nError: %s", err.Error())
	}

	a.DB = client
	a.Router = mux.NewRouter()
	a.setRoutes()
}

func (a *App) setRoutes() {
	// Account endpoints
	a.Get("/accounts", a.handleRequest(handler.GetAccounts))
	a.Get("/accounts/{account_id}/balance", a.handleRequest(handler.GetAccountBalance))
	a.Post("/accounts", a.handleRequest(handler.CreateAccount))

	// Transfer endpoints
	a.Get("/transfers", a.handleRequest(handler.GetTransfers))
	a.Post("/transfers", a.handleRequest(handler.MakeTransfer))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *mongo.Client, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}