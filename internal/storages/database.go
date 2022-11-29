package storages

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	createTable = `
	CREATE TABLE IF NOT EXISTS %s (
		userId TEXT,
		userName TEXT,
		name TEXT PRIMARY KEY, 
		birthday TEXT
		draft INTEGER
	);
	`
)

func NewDatabase(path string) error {
	if path == "" {
		return errors.New("invalid path")
	}

	_, err := sql.Open("sqlite3", path)
	if err != nil {
		return errors.Wrap(err, "creating a database")
	}

	return nil
}

func NewTable(path, name string) error {
	if path == "" {
		return errors.New("invalid path")
	}
	if name == "" {
		return errors.New("A table name is empty")
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return errors.Wrap(err, "opening the database")
	}

	cmd := fmt.Sprintf(createTable, name)

	if _, err = db.Exec(cmd); err != nil {
		return errors.Wrap(err, "creating a table with a name: "+name)
	}

	return nil
}
