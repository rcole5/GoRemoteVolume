package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
	// "github.com/rs/cors"
)

type Vol struct {
	Volume int  `json:"volume"`
	Muted  bool `json:"muted"`
}

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   Vol    `json:"data"`
}

// Returns the current volume.
func getCurrentVolume() (vol int) {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatal("ERR: Cannot get volume")
	}

	return vol
}

// Returns the current volume
// Value from 0-100
func getVolume(w http.ResponseWriter, r *http.Request) {
	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: false}}
	json.NewEncoder(w).Encode(volObj)
}

// Sets volume to the endpoint value
// Values from 0-100
func setVolume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newVol, err := strconv.Atoi(params["vol"])
	err = volume.SetVolume(newVol)
	if err != nil {
		volObj := Response{Status: 500, Error: "Could not set volume."}
		json.NewEncoder(w).Encode(volObj)
		return
	}

	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: false}}
	json.NewEncoder(w).Encode(volObj)
}

// Mute the volume if it is unmuted
// Unmutes if it is muted
func muteVolume(w http.ResponseWriter, r *http.Request) {
	isMute, err := volume.GetMuted()
	if err != nil {
		volObj := Response{Status: 500, Error: "Could not detect mute."}
		json.NewEncoder(w).Encode(volObj)
		return
	}

	if isMute {
		err = volume.Unmute()
		if err != nil {
			volObj := Response{Status: 500, Error: "Could not unmute."}
			json.NewEncoder(w).Encode(volObj)
			return
		}
	} else {

		err = volume.Mute()
		if err != nil {
			volObj := Response{Status: 500, Error: "Could not mute."}
			json.NewEncoder(w).Encode(volObj)
			return
		}
	}

	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: !isMute}}
	json.NewEncoder(w).Encode(volObj)
}

func main() {
	// Create new Router
	router := mux.NewRouter()

	// Set routes
	router.HandleFunc("/", getVolume).Methods("GET")
	router.HandleFunc("/volume/{vol}", setVolume).Methods("GET")
	router.HandleFunc("/mute", muteVolume).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Start server
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
