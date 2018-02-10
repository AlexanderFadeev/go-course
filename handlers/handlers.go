package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", list).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", video).Methods(http.MethodGet)
	s.HandleFunc("/video", postVideo).Methods(http.MethodPost)

	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.WithFields(log.Fields{
				"method":     r.Method,
				"url":        r.RequestURI,
				"remoteAddr": r.RemoteAddr,
				"userAgent":  r.UserAgent(),
			}).Info("New request")
			h.ServeHTTP(w, r)
		},
	)
}
