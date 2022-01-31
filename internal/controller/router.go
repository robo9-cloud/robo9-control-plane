package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func StartServer() {
	tc := NewTaskController()
	app := http.NewServeMux()
	app.Handle("/tasks", tc)
	app.Handle("/tasks/", tc)
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
