package handlers

import (
	"net/http"

	"github.com/AlexanderFadeev/go-course/database"
	"github.com/AlexanderFadeev/go-course/uploader"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	*mux.Router
}

func NewRouter(db database.DB, uploader uploader.Uploader) http.Handler {
	muxRouter := mux.NewRouter()
	r := Router{
		Router: muxRouter,
	}

	s := muxRouter.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", list(db)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", video(db)).Methods(http.MethodGet)
	s.HandleFunc("/video", postVideo(db, uploader)).Methods(http.MethodPost)

	return logMiddleware(&r)
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

func generateUUID() (*string, error) {
	idStruct, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to generate key")
	}
	id := idStruct.String()
	return &id, nil
}
