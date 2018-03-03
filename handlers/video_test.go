package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"path"

	"github.com/AlexanderFadeev/go-course/database"
)

func TestVideo(t *testing.T) {
	db, err := database.New("root", "1234", "video")
	if err != nil {
		t.Fatal(err)
	}

	videos, err := db.GetList()
	if err != nil {
		t.Fatal(err)
	}

	target := path.Join("/api/v1/", videos[0].URL)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, target, nil)
	video(db)(w, r)

	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	jsonData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var video database.Video
	err = json.Unmarshal(jsonData, &video)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON data: %v", err)
	}
}
