package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func StartServer() {
	tc := NewTaskController()

	app := mux.NewRouter()

	app.HandleFunc("/tasks", tc.GetAll).Methods("GET")
	app.HandleFunc("/tasks", tc.Post).Methods("POST")
	app.HandleFunc("/tasks/{id}", tc.Get).Methods("GET")
	app.HandleFunc("/tasks/{id}", tc.Put).Methods("PUT")
	app.HandleFunc("/tasks/{id}", tc.Delete).Methods("DELETE")

	s := &http.Server{
		Addr:    ":18080",
		Handler: app,
	}
	log.Fatal(s.ListenAndServe())
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
