package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlexanderFadeev/go-course/database"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func video(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["ID"]

		videoInfo, err := db.GetVideo(id)
		if err != nil {
			err = errors.Wrap(err, "Failed to get video info")
			log.Panic(err)
		}

		bytes, err := json.Marshal(videoInfo)
		if err != nil {
			err = errors.Wrap(err, "Failed to unmarshal to JSON")
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		_, err = io.WriteString(w, string(bytes))
		if err != nil {
			err = errors.Wrap(err, "Failed to write")
			log.Error(err)
		}
	}
}
