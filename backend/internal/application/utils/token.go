package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// Fonction pour générer un token sécurisé
func GenerateRandomToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
