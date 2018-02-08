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

func getVideoInfo(id string) (*videoInfo, error) {
	switch id {
	case "d290f1ee-6c54-4b01-90e6-d701748f0851":
		return &videoInfo{
			ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
			Name:      "Black Retrospetive Woman",
			Duration:  15,
			Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
			URL:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
		}, nil
	case "sldjfl34-dfgj-523k-jk34-5jk3j45klj34":
		return &videoInfo{
			ID:        "sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
			Name:      "Go Rally TEASER-HD",
			Duration:  41,
			Thumbnail: "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
			URL:       "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4",
		}, nil
	case "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345":
		return &videoInfo{
			ID:        "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
			Name:      "Танцор",
			Duration:  92,
			Thumbnail: "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
			URL:       "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4",
		}, nil
	default:
		return nil, errors.Errorf("Unknown video info %s", id)
	}

}

func video(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	videoInfo, err := getVideoInfo(id)
	if err != nil {
		err = errors.Wrap(err, "Failed to unmarshal to JSON")
		log.Error(err)
	}

	bytes, err := json.Marshal(videoInfo)
	if err != nil {
		err = errors.Wrap(err, "Failed to get video info")
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
