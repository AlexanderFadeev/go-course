package handlers

import (
	"encoding/json"
	"io"
	"net/http"

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

func getVideoInfo() *videoInfo {
	return &videoInfo{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		URL:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}
}

func video(w http.ResponseWriter, _ *http.Request) {
	l := getVideoInfo()

	bytes, err := json.Marshal(l)
	if err != nil {
		err = errors.Wrap(err, "Failed to unmarshal to JSON")
		log.Error(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_, err = io.WriteString(w, string(bytes))
	if err != nil {
		err = errors.Wrap(err, "Failed to write")
		log.Error(err)
	}
}
