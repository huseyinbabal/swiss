package gzip

import (
	"bytes"
	gz "compress/gzip"
	"encoding/base64"
	"io"
)

// Compress gzip-compresses the input bytes and returns the result as a
// standard base64-encoded string.
func Compress(in string) (string, error) {
	var buf bytes.Buffer
	w := gz.NewWriter(&buf)
	if _, err := w.Write([]byte(in)); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// Decompress base64-decodes the input then gunzips it, returning the
// original text.
func Decompress(in string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "", err
	}
	r, err := gz.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer r.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
