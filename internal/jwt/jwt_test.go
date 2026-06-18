package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
)

// sign builds an HS256 JWT from a header/payload/secret so the test never
// depends on a hand-computed signature.
func sign(header, payload, secret string) string {
	h := base64.RawURLEncoding.EncodeToString([]byte(header))
	p := base64.RawURLEncoding.EncodeToString([]byte(payload))
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(h + "." + p))
	s := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return h + "." + p + "." + s
}

const (
	sampleHeader  = `{"alg":"HS256","typ":"JWT"}`
	samplePayload = `{"sub":"1234567890","name":"John Doe","exp":2000000000}`
)

func sampleToken() string {
	return sign(sampleHeader, samplePayload, "secret")
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		wantErr   bool
		wantParts []string
	}{
		{
			name:      "valid token",
			token:     sampleToken(),
			wantErr:   false,
			wantParts: []string{"Header:", "Payload:", "John Doe", "Expires:"},
		},
		{
			name:    "wrong number of parts",
			token:   "a.b",
			wantErr: true,
		},
		{
			name:    "bad base64",
			token:   "!!!.!!!.!!!",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.token)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			for _, p := range tt.wantParts {
				if !strings.Contains(got, p) {
					t.Errorf("Decode() = %q, missing %q", got, p)
				}
			}
		})
	}
}

func TestVerify(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		secret  string
		want    string
		wantErr bool
	}{
		{
			name:   "valid signature",
			token:  sampleToken(),
			secret: "secret",
			want:   "valid",
		},
		{
			name:   "invalid signature",
			token:  sampleToken(),
			secret: "wrong",
			want:   "invalid",
		},
		{
			name:    "malformed token",
			token:   "a.b",
			secret:  "secret",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.token, tt.secret)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Verify() = %q, want %q", got, tt.want)
			}
		})
	}
}
