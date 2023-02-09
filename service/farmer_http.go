package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type MsgResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func registerHandler(deps dependencies) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var farmer NewFarmer
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
		logrus.Info(farmer.Password)
		addedFarmer, err := deps.FarmService.Register(req.Context(), farmer)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		respBytes, _ := json.Marshal(addedFarmer)
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

func addMachineHandler(deps dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authentication")
		if tokenString == "" {
			http.Error(w, "No token provided", http.StatusBadRequest)
			return
		}
		var machine NewMachine

		if err := json.NewDecoder(r.Body).Decode(&machine); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var err error
		machine.OwnerId, err = ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token provided", http.StatusBadRequest)
			return
		}
		addedMachine, err := deps.FarmService.AddMachine(r.Context(), machine)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		respBytes, _ := json.Marshal(addedMachine)
		w.Write(respBytes)
		w.WriteHeader(http.StatusCreated)
	}
}

func getMachineHandler(deps dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authentication")
		if tokenString == "" {
			http.Error(w, "No token provided", http.StatusBadRequest)
			return
		}
		_, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token provided", http.StatusBadRequest)
			return
		}
		machines, err := deps.FarmService.GetMachines(r.Context())

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		respBytes, _ := json.Marshal(machines)
		w.Write(respBytes)
		w.WriteHeader(http.StatusOK)
	}
}

func bookingHandler(deps dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// tokenString := r.Header.Get("Authentication")
		// if tokenString == "" {
		// 	http.Error(w, "No token provided", http.StatusBadRequest)
		// 	return
		// }
		// _, err := ValidateJWT(tokenString)
		// if err != nil {
		// 	http.Error(w, "Invalid token provided", http.StatusBadRequest)
		// 	return
		// }
		var booking NewBooking

		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := ValidateBookingslots(booking.Slots); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ValidateBookingDate(booking.Date); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		addedBooking, err := deps.FarmService.BookMachine(r.Context(), booking)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		respBytes, _ := json.Marshal(addedBooking)
		w.Write(respBytes)
		w.WriteHeader(http.StatusCreated)
	}
}

func availabilityHandler(deps dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		machineId := vars["id"]
		id, err := strconv.Atoi(machineId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		slotsAvailable, err := deps.FarmService.GetAvailability(r.Context(), uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		respBytes, _ := json.Marshal(slotsAvailable)
		w.Write(respBytes)
		w.WriteHeader(http.StatusCreated)
	}
}
