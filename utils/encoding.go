package utils

import (
	"encoding/hex"
)

// Converts a byte array to a string of hex encoding
func ByteArrayToHex(byteArray []byte) string {
	return hex.EncodeToString(byteArray)
}

// Converts a string of hex encoding to a byte array
func HexToByteArray(string string) ([]byte, error) {
	return hex.DecodeString(string)
}

// Converts a byte array to a string
func ByteArrayToString(byteArray []byte) string {
	return string(byteArray)
}

// Converts a string to a byte array
func StringToByteArray(string string) []byte {
	return []byte(string)
}
