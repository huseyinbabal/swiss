package qr

import (
	"bytes"
	"strings"
	"testing"
)

func TestTerminal(t *testing.T) {
	out, err := Terminal("https://github.com/huseyinbabal/swiss")
	if err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatal("expected non-empty output")
	}
	if !strings.ContainsAny(out, "█▀▄") {
		t.Error("output does not look like a block QR code")
	}
}

func TestTerminalEmpty(t *testing.T) {
	if _, err := Terminal(""); err == nil {
		t.Error("expected an error for empty input")
	}
}

func TestPNG(t *testing.T) {
	png, err := PNG("hello", 128)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(png, []byte("\x89PNG\r\n\x1a\n")) {
		t.Error("output is not a valid PNG")
	}
}

func TestPNGEmpty(t *testing.T) {
	if _, err := PNG("", 128); err == nil {
		t.Error("expected an error for empty input")
	}
}
