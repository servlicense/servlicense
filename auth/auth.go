package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Intevel/servlicense.sh/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

const (
	time    = 1
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
)

func GenerateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	return bytes, err
}

func CreateApiKey(name string, scopes []string) error {
	// api key is uuid
	apiKey := uuid.New().String()

	// hash the api key
	hash, err := HashApiKey(apiKey)
	if err != nil {
		return err
	}

	// insert the api key into the database
	err = database.Get().InsertApiKey(hash, name, scopes)

	return err
}

func HashApiKey(apiKey string) (string, error) {
	salt, err := GenerateRandomBytes(16)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(apiKey), salt, time, memory, threads, keyLen)

	// Encode the hash and salt into a single string.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("%s:%s", b64Salt, b64Hash)

	return encodedHash, nil
}

func VerifyApiKey(apiKey, encodedHash string) (bool, error) {
	// Split the encoded hash into the salt and hash.
	parts := strings.Split(encodedHash, ":")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid encoded hash")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	// Hash the apiKey with the salt.
	newHash := argon2.IDKey([]byte(apiKey), salt, time, memory, threads, keyLen)

	// Compare the hashes.
	if !bytes.Equal(hash, newHash) {
		return false, nil
	}

	return true, nil
}
