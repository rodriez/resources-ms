package app

import (
	"net/http"
	"resources-ms/handlers"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type myRouter struct {
	*mux.Router
}

var Router *myRouter

func init() {
	Router = &myRouter{mux.NewRouter()}
	Router.Use(corsMiddleware)
	Router.HandleFunc("/ping", handlers.Ping).Methods("GET")

	resRouter := Router.PathPrefix("/resource").Subrouter()
	resRouter.Use(handlers.BasicAuthMiddleware)
	resRouter.HandleFunc("/", handlers.CreateResource).Methods(http.MethodPost)
	resRouter.HandleFunc("/{id}", handlers.FindResource).Methods(http.MethodGet)
	resRouter.HandleFunc("/{id}", handlers.UpdateResource).Methods(http.MethodPut)
	resRouter.HandleFunc("/{id}", handlers.DeleteResource).Methods(http.MethodDelete)

	logrus.Info("Router initialized")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("Allow-Control-Allow-Origin", "*")
			res.Header().Set("Allow-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTION")
			res.Header().Set("Allow-Control-Allow-Headers", "Origin,Content-Type")

			if req.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(res, req)
		},
	)
}
