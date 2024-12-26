package licenses

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/servlicense/servlicense/api/database"
	"github.com/servlicense/servlicense/api/models"
)

func GetLicense(license string) (models.License, error) {
	licenseData, err := database.Get().GetLicense(license)
	if err != nil {
		return models.License{}, fmt.Errorf("failed to get license from database: %w", err)
	}

	return licenseData, nil
}

func CreateLicense(validUntil string) (uuid.UUID, error) {
	licenseUUID := uuid.New()

	createdAt := time.Now().Format(time.RFC3339)

	// Create the License struct
	license := models.License{
		License:    licenseUUID.String(),
		Active:     true,
		ValidUntil: validUntil,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}

	err := database.Get().InsertLicense(license)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to insert license into database: %w", err)
	}

	// Return the generated UUID
	return licenseUUID, nil
}

func UpdateLicense(license models.License) error {
	updatedAt := time.Now().Format(time.RFC3339)

	license.UpdatedAt = updatedAt

	err := database.Get().UpdateLicense(license)
	if err != nil {
		return fmt.Errorf("failed to update license in database: %w", err)
	}

	return nil
}

func CheckLicense(license string) (bool, error) {
	licenseData, err := GetLicense(license)
	if err != nil {
		return false, fmt.Errorf("failed to get license: %w", err)
	}

	// Check if the license is active
	if !licenseData.Active {
		return false, nil
	}

	// Check if the license is still valid
	validUntil, err := time.Parse(time.RFC3339, licenseData.ValidUntil)
	if err != nil {
		return false, fmt.Errorf("failed to parse valid until date: %w", err)
	}

	if time.Now().After(validUntil) {
		return false, nil
	}

	return true, nil
}

func ListLicenses(appID string) ([]models.License, error) {
	licenses, err := database.Get().ListLicenses(appID)
	if err != nil {
		return nil, fmt.Errorf("failed to list licenses: %w", err)
	}

	return licenses, nil
}
