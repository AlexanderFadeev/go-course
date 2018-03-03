package database

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type DB interface {
	GetList() ([]*Video, error)
	GetVideo(id string) (*Video, error)
	AddVideo(*Video) error
}

type impl struct {
	db   *sql.DB
	name string
}

func New(user, pass, name string) (DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, pass, name))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open sql database")
	}

	return &impl{
		db:   db,
		name: name,
	}, nil
}
