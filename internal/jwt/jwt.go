package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

func decodeSegment(seg string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(strings.TrimRight(seg, "="))
}

func prettyJSON(data []byte) (string, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return "", err
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Decode splits a JWT, base64url-decodes the header and payload, and
// pretty-prints them as indented JSON. If the payload has a numeric "exp"
// claim, the expiration time is appended in RFC3339 format.
func Decode(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token: expected 3 parts")
	}

	headerRaw, err := decodeSegment(parts[0])
	if err != nil {
		return "", err
	}
	payloadRaw, err := decodeSegment(parts[1])
	if err != nil {
		return "", err
	}

	headerJSON, err := prettyJSON(headerRaw)
	if err != nil {
		return "", err
	}
	payloadJSON, err := prettyJSON(payloadRaw)
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Header:\n%s\n\nPayload:\n%s", headerJSON, payloadJSON)

	var claims map[string]interface{}
	if err := json.Unmarshal(payloadRaw, &claims); err == nil {
		if exp, ok := claims["exp"]; ok {
			if expF, ok := exp.(float64); ok {
				t := time.Unix(int64(expF), 0).UTC()
				res += fmt.Sprintf("\n\nExpires: %s", t.Format(time.RFC3339))
			}
		}
	}

	return res, nil
}

// Verify checks the HS256 signature of a JWT against the given secret.
// It returns "valid" or "invalid"; an error is only returned for a
// malformed token.
func Verify(token, secret string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token: expected 3 parts")
	}

	sig, err := decodeSegment(parts[2])
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(parts[0] + "." + parts[1]))
	expected := mac.Sum(nil)

	if hmac.Equal(sig, expected) {
		return "valid", nil
	}
	return "invalid", nil
}
