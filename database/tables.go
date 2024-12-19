package database

import (
	"fmt"
	"log/slog"
)

var tableCreationSql []string = []string{
	"PRAGMA foreign_keys = ON;",
	"CREATE TABLE IF NOT EXISTS license (license TEXT PRIMARY KEY, active BOOLEAN, valid_until TEXT, created_at TEXT, updated_at TEXT)",
	"CREATE TABLE IF NOT EXISTS api_keys (id SERIAL PRIMARY KEY, api_key TEXT NOT NULL UNIQUE, name VARCHAR(255) NOT NULL, scopes TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);",
}

func (d *Database) CreateTablesIfNotExist() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	t, err := d.Db.Begin()
	if err != nil {
		return fmt.Errorf("Failed to begin transaction: %w", err)
	}
	for _, s := range tableCreationSql {
		slog.Debug("Attempting to create table", "sql", s)
		_, err := t.Exec(s)
		if err != nil {
			fmt.Println("Failed to create table:", err, "sql:", s)
			slog.Info("Rolling back the transaction")
			if err := t.Rollback(); err != nil {
				return fmt.Errorf("Failed to roll back transaction: %w", err)
			}
			return err
		}
	}
	err = t.Commit()
	if err != nil {
		return fmt.Errorf("Failed to commit transaction: %w", err)
	}

	slog.Debug("Successfully created tables")
	return nil
}
