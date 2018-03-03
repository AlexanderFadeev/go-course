package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexanderFadeev/go-course/database"
)

func TestList(t *testing.T) {
	db, err := database.New("root", "1234", "video_test")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	list(db)(w, nil)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	jsonData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var videos []*database.Video
	err = json.Unmarshal(jsonData, &videos)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON data: %v", err)
	}
}
