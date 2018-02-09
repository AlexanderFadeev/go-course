package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func TestList(t *testing.T) {
	w := httptest.NewRecorder()
	list(w, nil)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	jsonData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	items := []videoListItem{}
	err = json.Unmarshal(jsonData, &items)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON data: %v", err)
	}
}
