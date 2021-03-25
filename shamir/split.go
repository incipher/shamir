package shamir

import (
	"fmt"

	"github.com/hashicorp/vault/shamir"
	"incipher.io/shamir/utils"
)

// Splits a secret into shares (of length sharesCount),
// which a subset thereof (of length thresholdCount) is
// necessary to reconstruct the original secret.
func Split(
	secret string,
	sharesCount int,
	thresholdCount int,
) ([]string, error) {
	// Validate inputs
	if len(secret) == 0 {
		return nil, fmt.Errorf("secret cannot be empty")
	}

	if thresholdCount > sharesCount {
		return nil, fmt.Errorf("threshold must be less than or equal shares")
	}

	if sharesCount < 2 || sharesCount > 255 {
		return nil, fmt.Errorf("shares must be between 2 and 255")
	}

	if thresholdCount < 2 || thresholdCount > 255 {
		return nil, fmt.Errorf("threshold must be between 2 and 255")
	}

	// Split secret into shares
	sharesBytes, err := shamir.Split(
		utils.StringToByteArray(secret),
		sharesCount,
		thresholdCount,
	)

	if err != nil {
		return nil, err
	}

	// Encode shares in hexadecimal
	sharesHex := make([]string, len(sharesBytes))

	for i := range sharesBytes {
		sharesHex[i] = utils.ByteArrayToHex(sharesBytes[i])
	}

	// Return shares
	return sharesHex, nil
}
