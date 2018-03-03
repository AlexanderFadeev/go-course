package handlers

import (
	"fmt"
	"net/http"

	"github.com/AlexanderFadeev/go-course/uploader"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func (r *Router) postVideo(uploader uploader.Uploader) http.HandlerFunc {
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

		videoKey, err := uuid.NewV4()
		if err != nil {
			err = errors.Wrap(err, "Failed to generate key")
			logrus.Panic(err)
		}
		const defaultTitle = "Title"
		const defaultStatus = 3
		const defaultDuration = 42

		url := fmt.Sprintf("/content/%s/index.mp4", videoKey)
		thumbnailURL := "/default_thumbnail.jpg"

		q := `
		INSERT INTO video SET
		video_key = ?,
		title = ?,
		status = ?,
		duration = ?,
		url = ?,
		thumbnail_url = ?
		`

		r.db.Query(q,
			videoKey,
			defaultTitle,
			defaultStatus,
			defaultDuration,
			url,
			thumbnailURL,
		)

		err = uploader.Upload(srcFile, url)
		if err != nil {
			err = errors.Wrap(err, "Failed to upload file")
			logrus.Panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
