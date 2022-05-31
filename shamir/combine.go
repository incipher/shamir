package shamir

import (
	"fmt"

	"github.com/hashicorp/vault/shamir"
	"incipher.io/shamir/utils"
)

// Reconstructs a secret from shares.
func Combine(sharesHex []string) (string, error) {
	// Convert shares from hexadecimal to byte arrays
	sharesBytes := make([][]byte, len(sharesHex))

	for i := range sharesHex {
		shareBytes, err := utils.HexToByteArray(sharesHex[i])
		if err != nil {
			return "", err
		}

		sharesBytes[i] = shareBytes
	}

	// Validate inputs
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

	// Reconstruct secret from shares
	secretBytes, err := shamir.Combine(sharesBytes)
	if err != nil {
		return "", err
	}

	// Convert secret from byte array to string
	secret := utils.ByteArrayToString(secretBytes)

	// Return secret
	return secret, nil
}
