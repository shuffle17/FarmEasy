package service

import (
	"FarmEasy/db"
	"encoding/json"
	"net/http"
)

type MsgResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func registerHandler(deps dependencies) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var farmer db.Farmer
		err := json.NewDecoder(req.Body).Decode(&farmer)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = ValidateFarmerPhone(farmer.Phone); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		if err = ValidateFarmerEmail(farmer.Email); err != nil {

			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		farmer, err = deps.FarmService.Register(req.Context(), farmer)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		respBytes, _ := json.Marshal(farmer)
		rw.Write(respBytes)
		rw.WriteHeader(http.StatusCreated)
	})
}

func loginHandler(deps dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fAuth LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&fAuth); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := ValidateFarmerEmail(fAuth.Email); err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tokenString, err := deps.FarmService.Login(r.Context(), fAuth)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		msg := MsgResponse{Message: "Login Successful", Token: tokenString}
		respBytes, _ := json.Marshal(msg)
		w.Write(respBytes)
		w.WriteHeader(http.StatusOK)
	}
}
