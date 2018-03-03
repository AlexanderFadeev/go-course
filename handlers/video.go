package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type videoInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
}

func (r *Router) getVideoInfo(id string) (*videoInfo, error) {
	q := `
		SELECT video_key, title, duration, url, thumbnail_url 
		FROM video 
		WHERE video_key = ?
		`

	rows, err := r.db.Query(q, id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send query to database")
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Unknown video ID")
	}

	var info videoInfo
	rows.Scan(
		&info.ID,
		&info.Name,
		&info.Duration,
		&info.URL,
		&info.Thumbnail,
	)
	return &info, nil
}

func (r *Router) video(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["ID"]

	videoInfo, err := r.getVideoInfo(id)
	if err != nil {
		err = errors.Wrap(err, "Failed to unmarshal to JSON")
		log.Panic(err)
	}

	bytes, err := json.Marshal(videoInfo)
	if err != nil {
		err = errors.Wrap(err, "Failed to get video info")
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
