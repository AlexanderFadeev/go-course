package handlers

import (
	"fmt"
	"net/http"

	"github.com/AlexanderFadeev/go-course/database"
	"github.com/AlexanderFadeev/go-course/uploader"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func postVideo(db database.DB, uploader uploader.Uploader) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		const key = "file[]"
		srcFile, fileHeader, err := req.FormFile(key)
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

		id, err := generateUUID()
		if err != nil {
			err = errors.Wrap(err, "Failed to generate video ID")
			logrus.Panic(err)
		}

		url := fmt.Sprintf("/content/%s/index.mp4", *id)

		video := database.Video{
			ID:   *id,
			URL:  url,
			Name: filename,
		}

		err = db.AddVideo(&video)
		if err != nil {
			err = errors.Wrap(err, "Failed to add video to database")
			logrus.Panic(err)
		}

		err = uploader.Upload(srcFile, url)
		if err != nil {
			err = errors.Wrap(err, "Failed to upload file")
			logrus.Panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
