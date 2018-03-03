package database

import (
	"fmt"

	"github.com/pkg/errors"
)

func (i *impl) GetList() ([]*Video, error) {
	query := fmt.Sprintf(`
		SELECT video_key, title, duration, thumbnail_url
		FROM %s;
		`, i.name)

	rows, err := i.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send query to database")
	}
	defer rows.Close()

	var list []*Video
	for rows.Next() {
		var video Video
		err := rows.Scan(
			&video.ID,
			&video.Name,
			&video.Duration,
			&video.Thumbnail,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to scan list items from rows")
		}

		list = append(list, &video)
	}

	return list, nil
}
