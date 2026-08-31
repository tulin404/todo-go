package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateID() (string, error) {
	id := make([]byte, 8)

	for i := range id {
		number, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", fmt.Errorf("failed to generate uuid: %v", err)
		}

		id[i] = alphabet[number.Int64()]
	}

	return string(id), nil
}
