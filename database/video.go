package database

import (
	"fmt"

	"github.com/pkg/errors"
)

type Video struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	Status    int    `json:"status"`
}

func (i *impl) AddVideo(videoOrig *Video) error {
	const defaultName = "Title"
	const defaultDuration = 42
	const defaultThumbnail = "/default_thumbnail.jpg"
	const defaultStatus = 3

	video := *videoOrig
	if video.URL == "" {
		return errors.New("URL is not provided")
	}

	if video.ID == "" {
		return errors.New("ID is not provided")
	}

	if video.Name == "" {
		video.Name = defaultName
	}

	if video.Duration == 0 {
		video.Duration = defaultDuration
	}

	if video.Thumbnail == "" {
		video.Thumbnail = defaultThumbnail
	}

	if video.Status == 0 {
		video.Status = defaultStatus
	}

	return i.addVideoImpl(&video)

}

func (i *impl) GetVideo(id string) (*Video, error) {
	query := fmt.Sprintf(`
		SELECT video_key, title, duration, url, thumbnail_url 
		FROM %s
		WHERE video_key = ?
		`, i.name)

	rows, err := i.db.Query(query, id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send query to database")
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Unknown video ID")
	}

	video := Video{}
	err = rows.Scan(
		&video.ID,
		&video.Name,
		&video.Duration,
		&video.URL,
		&video.Thumbnail,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to scan rows")
	}

	return &video, nil
}

func (i *impl) addVideoImpl(video *Video) error {
	query := fmt.Sprintf(`
		INSERT INTO %s SET
		video_key = ?,
		title = ?,
		status = ?,
		duration = ?,
		url = ?,
		thumbnail_url = ?
		`, i.name)

	rows, err := i.db.Query(query,
		&video.ID,
		&video.Name,
		&video.Status,
		&video.Duration,
		&video.URL,
		&video.Thumbnail,
	)
	if err != nil {
		return errors.Wrap(err, "Failed to send query to database")
	}

	rows.Close()
	return nil
}
