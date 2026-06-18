package qr

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

// Terminal renders content as a compact, scannable QR code using half-block
// characters, suitable for printing directly in a terminal.
func Terminal(content string) (string, error) {
	if content == "" {
		return "", fmt.Errorf("cannot encode an empty value")
	}
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return "", err
	}
	return q.ToSmallString(false), nil
}

// PNG returns a PNG image of the QR code for content at the given pixel size.
func PNG(content string, size int) ([]byte, error) {
	if content == "" {
		return nil, fmt.Errorf("cannot encode an empty value")
	}
	if size <= 0 {
		size = 256
	}
	return qrcode.Encode(content, qrcode.Medium, size)
}
