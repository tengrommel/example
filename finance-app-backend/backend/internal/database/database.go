package database

import (
	"github.com/jmoiron/sqlx"
	"io"
)

type Database interface {
	io.Closer
}

type database struct {
	conn *sqlx.DB
}

func (d *database) Close() error {
	return d.conn.Close()
}
