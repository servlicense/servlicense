package database

import (
	"strings"

	"github.com/Intevel/servlicense.sh/models"
)

func (d *Database) InsertApiKey(id string, apiKey string, name string, scopes []string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	scopesStr := strings.Join(scopes, ",")
	_, err := d.Db.Exec("INSERT INTO api_keys (id, api_key, name, scopes) VALUES(?, ?, ?, ?)", id, apiKey, name, scopesStr)
	return err
}

func (d *Database) GetApiKey(id string) (models.Apikey, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	row := d.Db.QueryRow("SELECT * FROM api_keys WHERE id = ?", id)
	var apikey models.Apikey
	err := row.Scan(&apikey.Id, &apikey.ApiKey, &apikey.Name, &apikey.Scopes, &apikey.CreatedAt)
	return apikey, err
}
