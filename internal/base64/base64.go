package base64

import "encoding/base64"

func Encode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func Decode(in string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "", err
	}
	return string(decodeString), nil
}
