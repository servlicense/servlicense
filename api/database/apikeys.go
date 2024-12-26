package database

import (
	"strings"

	"github.com/servlicense/servlicense/api/models"
)

func (d *Database) InsertApiKey(id string, apiKey string, name string, scopes []string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	scopesStr := strings.Join(scopes, ",")
	_, err := d.Db.Exec("INSERT INTO api_keys (id, api_key, name, scopes) VALUES(?, ?, ?, ?)", id, apiKey, name, scopesStr)
	return err
}

func (d *Database) GetApiKey(id string) (models.ApiKey, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	row := d.Db.QueryRow("SELECT * FROM api_keys WHERE id = ?", id)
	var apikey models.ApiKey
	var scopesStr string
	err := row.Scan(&apikey.Id, &apikey.ApiKey, &apikey.Name, &scopesStr, &apikey.CreatedAt)
	if err == nil {
		apikey.Scopes = strings.Split(scopesStr, ",")
	}
	return apikey, err
}

func (d *Database) ListApiKeys() ([]models.ApiKey, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	rows, err := d.Db.Query("SELECT id, name, scopes, created_at FROM api_keys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var apikeys []models.ApiKey
	for rows.Next() {
		var apikey models.ApiKey
		var scopesStr string
		err := rows.Scan(&apikey.Id, &apikey.Name, &scopesStr, &apikey.CreatedAt)
		if err != nil {
			return nil, err
		}
		apikey.Scopes = strings.Split(scopesStr, ",")
		apikeys = append(apikeys, apikey)
	}
	return apikeys, nil
}
