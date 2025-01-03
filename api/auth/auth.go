package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/godruoyi/go-snowflake"
	"github.com/google/uuid"
	"github.com/servlicense/servlicense/api/database"
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

func CreateApiKey(name string, scopes []string) (string, string, error) {
	// identifier for the api key
	id := snowflake.ID()
	idStr := strconv.FormatUint(id, 10)
	// api key is uuid
	apiKey := uuid.New().String()

	// hash the api key
	hash, err := HashApiKey(apiKey)
	if err != nil {
		return "", "", err
	}

	// insert the api key into the database, TODO replace fmt with strconv
	err = database.Get().InsertApiKey(idStr, hash, name, scopes)

	return idStr, apiKey, err
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

	// this could also be str+':'+str, but i think this should be faster
	b := strings.Builder{}
	b.WriteString(b64Salt)
	b.WriteRune(':')
	b.WriteString(b64Hash)
	encodedHash := b.String()

	fmt.Println(encodedHash)
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
