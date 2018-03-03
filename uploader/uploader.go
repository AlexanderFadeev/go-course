package uploader

import (
	"io"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Uploader interface {
	Upload(io.Reader, string) error
}

type impl struct {
	staticDir string
}

func New(staticDir string) Uploader {
	return &impl{staticDir: staticDir}
}

func (i *impl) Upload(src io.Reader, url string) error {
	dstPath := path.Join(i.staticDir, url)
	err := os.MkdirAll(path.Dir(dstPath), os.ModePerm)
	if err != nil {
		err = errors.Wrap(err, "Failed to create directory for destination file")
		logrus.Panic(err)
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		err = errors.Wrap(err, "Failed to open destination file")
		logrus.Panic(err)
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		err = errors.Wrap(err, "Failed to copy file")
		logrus.Panic(err)
	}

	return nil
}
