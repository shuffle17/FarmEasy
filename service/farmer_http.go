package service

import (
	"FarmEasy/db"
	"encoding/json"
	"net/http"
)

type MsgResponse struct {
	Message string `json:"message"`
}

func registerHandler(deps dependencies) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var farmer db.Farmer
		err := json.NewDecoder(req.Body).Decode(&farmer)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = ValidateFarmerPhone(farmer); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		if err = ValidateFarmerEmail(farmer); err != nil {

			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		farmer, err = deps.FarmService.Register(req.Context(), farmer)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		// response := PingResponse{Message: "Farmer added successfully"}
		respBytes, _ := json.Marshal(farmer)
		rw.Write(respBytes)
		rw.WriteHeader(http.StatusCreated)
	})
}
