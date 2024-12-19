package database

func (d *Database) InsertApiKey(apiKey string, name string, scopes []string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, err := d.Db.Exec("INSERT INTO api_keys (api_key, name, scopes) VALUES(?, ?, ?)", apiKey, name, scopes)
	return err
}
