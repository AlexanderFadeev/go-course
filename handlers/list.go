package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type videoListItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

var videolist []videoListItem = []videoListItem{
	{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	},
	{
		ID:        "sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
		Name:      "Go Rally TEASER-HD",
		Duration:  41,
		Thumbnail: "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
	},
	{
		ID:        "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
		Name:      "Танцор",
		Duration:  92,
		Thumbnail: "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
	},
}

func list(w http.ResponseWriter, _ *http.Request) {
	bytes, err := json.Marshal(videolist)
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
