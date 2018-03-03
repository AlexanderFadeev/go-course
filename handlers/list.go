package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlexanderFadeev/go-course/database"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func list(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		items, err := db.GetList()
		if err != nil {
			err = errors.Wrap(err, "Failed to get items list")
			log.Panic(err)
		}

		bytes, err := json.Marshal(items)
		if err != nil {
			err = errors.Wrap(err, "Failed to unmarshal to JSON")
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		_, err = io.WriteString(w, string(bytes))
		if err != nil {
			err = errors.Wrap(err, "Failed to write")
			log.Panic(err)
		}
	}
}
