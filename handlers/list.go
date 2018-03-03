package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type videoListItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

func (r *Router) getItemsList() ([]videoListItem, error) {
	q := `
		SELECT video_key, title, duration, thumbnail_url
		FROM video
		`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send query to database")
	}
	defer rows.Close()

	var items []videoListItem
	for rows.Next() {
		var item videoListItem
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Duration,
			&item.Thumbnail,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to scan list items from rows")
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *Router) list(w http.ResponseWriter, _ *http.Request) {
	items, err := r.getItemsList()
	if err != nil {
		err = errors.Wrap(err, "Failed to get items list")
		log.Panic(err)
	}

	bytes, err := json.Marshal(items)
	if err != nil {
		err = errors.Wrap(err, "Failed to unmarshal to JSON")
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_, err = io.WriteString(w, string(bytes))
	if err != nil {
		err = errors.Wrap(err, "Failed to write")
		log.Panic(err)
	}
}
