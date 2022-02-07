package controller

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		zap.S().Infof(r.Method + " " + r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func StartServer() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	tc := NewTaskController()

	app := mux.NewRouter()

	app.HandleFunc("/v1/tasks", tc.GetAll).Methods("GET")
	app.HandleFunc("/v1/tasks", tc.Post).Methods("POST")
	app.HandleFunc("/v1/tasks/{id}", tc.Get).Methods("GET")
	app.HandleFunc("/v1/tasks/{id}", tc.Put).Methods("PUT")
	app.HandleFunc("/v1/tasks/{id}", tc.Delete).Methods("DELETE")
	app.Use(loggingMiddleware)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      app,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	err := s.Shutdown(ctx)
	if err != nil {
		zap.S().Info(err.Error())
	}
	zap.S().Info("shutting down server")
	os.Exit(0)

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err != nil {
		zap.S().Info(err.Error())
	}
}
