package cert

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
	"time"
)

// Info parses a PEM-encoded X.509 certificate and returns a multi-line summary.
func Info(pemStr string) (string, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM data")
	}

	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse certificate: %w", err)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Subject: %s\n", c.Subject.String())
	fmt.Fprintf(&b, "Issuer: %s\n", c.Issuer.String())
	fmt.Fprintf(&b, "Serial: %s\n", c.SerialNumber.String())
	fmt.Fprintf(&b, "Not Before: %s\n", c.NotBefore.Format(time.RFC3339))
	fmt.Fprintf(&b, "Not After: %s\n", c.NotAfter.Format(time.RFC3339))
	fmt.Fprintf(&b, "DNS SANs: %s\n", strings.Join(c.DNSNames, ", "))
	fmt.Fprintf(&b, "IsCA: %t", c.IsCA)

	return b.String(), nil
}
