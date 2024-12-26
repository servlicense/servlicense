package database

import "github.com/servlicense/servlicense/api/models"

func (d *Database) GetLicense(license string) (models.License, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	row := d.Db.QueryRow("SELECT license, app_id, active, valid_until, created_at, updated_at FROM license WHERE license = ?", license)
	var l models.License
	err := row.Scan(&l.License, &l.AppID, &l.Active, &l.ValidUntil, &l.CreatedAt, &l.UpdatedAt)
	return l, err
}

func (d *Database) InsertLicense(license models.License) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, err := d.Db.Exec(
		"INSERT INTO license (license, app_id, active, valid_until, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)",
		license.License, license.AppID, license.Active, license.ValidUntil, license.CreatedAt, license.UpdatedAt,
	)
	return err
}

func (d *Database) UpdateLicense(license models.License) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, err := d.Db.Exec(
		"UPDATE license SET app_id = ?, active = ?, valid_until = ?, updated_at = ? WHERE license = ?",
		license.AppID, license.Active, license.ValidUntil, license.UpdatedAt, license.License,
	)
	return err
}

func (d *Database) ListLicenses(appID string) ([]models.License, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	rows, err := d.Db.Query("SELECT license, app_id, active, valid_until, created_at, updated_at FROM license WHERE app_id = ?", appID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var licenses []models.License
	for rows.Next() {
		var l models.License
		err := rows.Scan(&l.License, &l.AppID, &l.Active, &l.ValidUntil, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			return nil, err
		}
		licenses = append(licenses, l)
	}
	return licenses, nil
}
