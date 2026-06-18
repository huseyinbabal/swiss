package cert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey error: %v", err)
	}

	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1234),
		Subject:      pkix.Name{CommonName: "swiss.example.com"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		DNSNames:     []string{"swiss.example.com"},
		IsCA:         true,
	}

	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("CreateCertificate error: %v", err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})

	out, err := Info(string(pemBytes))
	if err != nil {
		t.Fatalf("Info error: %v", err)
	}
	if !strings.Contains(out, "swiss.example.com") {
		t.Fatalf("output does not contain subject CN: %q", out)
	}
}

func TestInfoInvalid(t *testing.T) {
	if _, err := Info("garbage"); err == nil {
		t.Fatalf("expected error for invalid PEM input, got nil")
	}
}
