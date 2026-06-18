package hex

import (
	stdhex "encoding/hex"
)

// Encode hex-encodes the input string. The error return is kept for
// uniformity and is always nil.
func Encode(in string) (string, error) {
	return stdhex.EncodeToString([]byte(in)), nil
}

// Decode hex-decodes the input string and returns the resulting text.
func Decode(in string) (string, error) {
	b, err := stdhex.DecodeString(in)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
