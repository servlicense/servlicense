package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	_ "modernc.org/sqlite"
)

var def *Database

func Get() *Database {
	if def == nil {
		def = &Database{}
	}
	return def
}

type Database struct {
	Db    *sql.DB
	mutex sync.Mutex
}

func (d *Database) Close() error {
	return d.Db.Close()
}

func (d *Database) Connect(path string) error {
	if path[0] != ':' {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Failed to get current directory: %w", err)
		}
		path = filepath.Join(currentDir, path)
		if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
			return err
		}
	}
	slog.Info("Connecting to database", "path", path)
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("Failed to establish database connection: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to ping database: %w", err)
	}
	d.Db = db
	return nil
}
