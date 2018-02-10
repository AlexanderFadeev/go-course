package handlers

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/pkg/errors"
)

func postVideo(w http.ResponseWriter, r *http.Request) {
	const key = "file[]"
	_, fileHeader, err := r.FormFile(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errors.Wrap(err, "Failed to get form file from request")
		logrus.Panic(err)
	}

	contentType := fileHeader.Header.Get("Content-Type")
	filename := fileHeader.Filename

	logrus.WithFields(logrus.Fields{
		"content_type": contentType,
		"filename":     filename,
	}).Info("Got file")

	const expectedType = "video/mp4"
	if contentType != expectedType {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
