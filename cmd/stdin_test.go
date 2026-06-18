package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestResolveStdinReadsValueFlag(t *testing.T) {
	c := &cobra.Command{}
	var v string
	c.Flags().StringVar(&v, "value", "", "")
	if err := c.Flags().Set("value", "-"); err != nil {
		t.Fatal(err)
	}
	c.SetIn(strings.NewReader("piped content"))

	if err := resolveStdin(c); err != nil {
		t.Fatalf("resolveStdin: %v", err)
	}
	if v != "piped content" {
		t.Errorf("value = %q, want %q", v, "piped content")
	}
}

func TestResolveStdinReadsUrlFlag(t *testing.T) {
	c := &cobra.Command{}
	var u string
	c.Flags().StringVar(&u, "url", "", "")
	_ = c.Flags().Set("url", "-")
	c.SetIn(strings.NewReader("a b&c"))

	if err := resolveStdin(c); err != nil {
		t.Fatal(err)
	}
	if u != "a b&c" {
		t.Errorf("url = %q, want %q", u, "a b&c")
	}
}

func TestResolveStdinLeavesLiteralValue(t *testing.T) {
	c := &cobra.Command{}
	var v string
	c.Flags().StringVar(&v, "value", "", "")
	_ = c.Flags().Set("value", "hello")
	c.SetIn(strings.NewReader("should be ignored"))

	if err := resolveStdin(c); err != nil {
		t.Fatal(err)
	}
	if v != "hello" {
		t.Errorf("value = %q, want %q (stdin must not be read)", v, "hello")
	}
}

func TestResolveStdinNoValueFlag(t *testing.T) {
	c := &cobra.Command{} // command without a value/url flag
	if err := resolveStdin(c); err != nil {
		t.Errorf("resolveStdin on flagless command: %v", err)
	}
}
