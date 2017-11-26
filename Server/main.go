package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
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

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")
	procKeyBd = moduser32.NewProc("keybd_event")
)

const (
	_KEYEVENTF_KEYUP = 0x0002
	_KEY_PLAY_PAUSE = 0xB3
	_KEY_TRACK_NEXT = 0xB0
	_KEY_TRACK_PREV = 0xB1
	_KEY_TRACK_STOP = 0xB2
)

// Returns the current volume.
func getCurrentVolume() (vol int) {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatal("ERR: Cannot get volume")
	}

	return vol
}

// Play a track if it's paused
// Pause a track if it's playing
func playPause(w http.ResponseWriter, r *http.Request) {
	sendKey(_KEY_PLAY_PAUSE)
	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: false}}
	json.NewEncoder(w).Encode(volObj)
}

// Play the next track
func nextTrack(w http.ResponseWriter, r *http.Request) {
	sendKey(_KEY_TRACK_NEXT)
	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: false}}
	json.NewEncoder(w).Encode(volObj)
}

// Play the previous track
func prevTrack(w http.ResponseWriter, r *http.Request) {
	sendKey(_KEY_TRACK_PREV)
	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: false}}
	json.NewEncoder(w).Encode(volObj)
}

func stopTrack(w http.ResponseWriter, r *http.Request) {
	sendKey(_KEY_TRACK_STOP)
	volObj := Response{Status: 200, Data: Vol{Volume: getCurrentVolume(), Muted: false}}
	json.NewEncoder(w).Encode(volObj)
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

// Send a key down and key up call to system
func sendKey(vk int) {
	downKey(vk)
	upKey(vk)
}

// Call key down to system
func downKey(key int) {
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), 0, 0)
}

// Call key up to system
func upKey(key int) {
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), _KEYEVENTF_KEYUP, 0)
}

func main() {
	// Create new Router
	router := mux.NewRouter()

	// Set routes
	router.HandleFunc("/", getVolume).Methods("GET")
	router.HandleFunc("/volume/{vol}", setVolume).Methods("GET")
	router.HandleFunc("/mute", muteVolume).Methods("GET")
	router.HandleFunc("/playpause", playPause).Methods("GET")
	router.HandleFunc("/next", nextTrack).Methods("GET")
	router.HandleFunc("/prev", prevTrack).Methods("GET")
	router.HandleFunc("/stop", stopTrack).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Start server
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
