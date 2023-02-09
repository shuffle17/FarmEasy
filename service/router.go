package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	router.HandleFunc("/register", registerHandler(deps)).Methods(http.MethodPost)

	router.HandleFunc("/login", loginHandler(deps)).Methods(http.MethodPost)

	router.HandleFunc("/machines", addMachineHandler(deps)).Methods(http.MethodPost)

	router.HandleFunc("/machines", getMachineHandler(deps)).Methods(http.MethodGet)

	router.HandleFunc("/bookings", bookingHandler(deps)).Methods(http.MethodPost)

	return
}
