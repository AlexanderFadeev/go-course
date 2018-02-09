package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideo(t *testing.T) {
	w := httptest.NewRecorder()
	const target = "/api/v1/video/d290f1ee-6c54-4b01-90e6-d701748f0851"
	req := httptest.NewRequest(http.MethodGet, target, nil)

	router := Router()
	router.ServeHTTP(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	jsonData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	info := videoInfo{}
	err = json.Unmarshal(jsonData, &info)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON data: %v", err)
	}
}
