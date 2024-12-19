package licenses

import (
	"fmt"
	"time"

	"github.com/Intevel/servlicense.sh/database"
	"github.com/Intevel/servlicense.sh/models"
	"github.com/google/uuid"
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
