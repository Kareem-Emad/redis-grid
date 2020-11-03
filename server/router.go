package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kareem-Emad/redis-grid/distributer"
	"github.com/gorilla/mux"
)

// listens to the assigned port for serving requests using the passed grpc server instance
func serve(router *mux.Router) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func route(router *mux.Router) {
	var sharder distributer.ShardsManager
	var response ResponsePayload

	sharder.InitShardsWithDefault()

	router.HandleFunc("/execute_command", func(w http.ResponseWriter, r *http.Request) {
		var payload CommandPayload

		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := sharder.ExecuteCommand(payload.Key, payload.CommandName, payload.CommandArgs...)

		response.CommandResult = result
		if err != nil {
			response.ErrorLog = err.Error()
		} else {
			response.ErrorLog = ""
		}

		log.Println("serving response ", response, " for request", payload)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("POST")
}

// Start maps all necessary routes and listens to the assigned port for serving requests
func Start() {
	router := mux.NewRouter()

	route(router)
	serve(router)
}
