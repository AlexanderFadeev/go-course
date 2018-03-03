package handlers

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"fmt"
	"os"
	"path"
	"io"
)

func (r *Router) postVideo(w http.ResponseWriter, req *http.Request) {
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

	dstFilePath := path.Join(r.staticDir, url)
	err = os.MkdirAll(path.Dir(dstFilePath), 0666)
	if err != nil {
		err = errors.Wrap(err, "Failed to create directory for destination file")
		logrus.Panic(err)
	}

	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to open destination file")
		logrus.Panic(err)
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		err = errors.Wrap(err, "Failed to copy file")
		logrus.Panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
