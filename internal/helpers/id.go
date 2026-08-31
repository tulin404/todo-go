package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 'GenerateID' generates and returns a 8 digit Base62 ID
func GenerateID() (string, error) {
	id := make([]byte, 8)

	for i := range id {
		number, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", fmt.Errorf("failed to generate uuid: %w", err)
		}

		id[i] = alphabet[number.Int64()]
	}

	return string(id), nil
}
