package handlers

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	impl      *mux.Router
	db        *sql.DB
	staticDir string
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.impl.ServeHTTP(w, req)
}

func (r *Router) HandleFunc(path string, handler http.HandlerFunc) *mux.Route {
	return r.impl.HandleFunc(path, handler)
}

func NewRouter(staticDir string) http.Handler {
	db, err := sql.Open("mysql", `root:1234@/video`)
	if err != nil {
		err = errors.Wrap(err, "Failed to open sql database")
		log.Panic(err)
	}

	muxRouter := mux.NewRouter()
	r := Router{
		impl:      muxRouter,
		db:        db,
		staticDir: staticDir,
	}

	s := muxRouter.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", r.list).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", r.video).Methods(http.MethodGet)
	s.HandleFunc("/video", r.postVideo).Methods(http.MethodPost)

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
