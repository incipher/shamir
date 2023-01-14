package shamir

import (
	"fmt"

	"github.com/hashicorp/vault/shamir"
	"incipher.io/shamir/utils"
)

// Reconstructs a secret from shares.
func Combine(sharesHex []string) (string, error) {
	sharesBytes := make([][]byte, len(sharesHex))

	for i := range sharesHex {
		shareBytes, err := utils.HexToByteArray(sharesHex[i])
		if err != nil {
			return "", err
		}

		sharesBytes[i] = shareBytes
	}

	if len(sharesBytes) < 2 || len(sharesBytes) > 255 {
		return "", fmt.Errorf("shares must be between 2 and 255")
	}

	firstShareBytes := sharesBytes[0]

	if len(firstShareBytes) < 2 {
		return "", fmt.Errorf("shares must be of length greater than 2 bytes")
	}

	for i := 1; i < len(sharesBytes); i++ {
		if len(sharesBytes[i]) != len(firstShareBytes) {
			return "", fmt.Errorf("shares must be of equal length")
		}
	}

	secretBytes, err := shamir.Combine(sharesBytes)
	if err != nil {
		return "", err
	}

	secret := utils.ByteArrayToString(secretBytes)

	return secret, nil
}
