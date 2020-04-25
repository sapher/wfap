package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ApiPort = 8080

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"messafe"`
	Payload interface{} `json:"payload"`
}

type ApiConfigureRequest struct {
	Password string `json:"password"`
	Ssid     string `json:"ssid"`
}

func main() {
	log.Println("wfap service is starting")

	// Handlers
	statusHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("status handler")
	}
	configureHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("configure handler")
	}

	// Build router
	router := mux.NewRouter()
	router.HandleFunc("/status", statusHandler)
	router.HandleFunc("/configure", configureHandler).Methods(http.MethodPost)
	http.Handle("/", router)

	// Server HTTP
	host := fmt.Sprintf("%s:%d", "", ApiPort)
	log.Printf("Server listenning on host: %s\n", host)
	err := http.ListenAndServe(host, handlers.CORS()(router))
	if err != nil {
		log.Fatal(err)
	}
}
