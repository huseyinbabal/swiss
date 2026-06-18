package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// stdinFlags are the value-bearing flags whose literal "-" means
// "read this value from stdin", enabling pipes such as:
//
//	echo "hello" | swiss base64 encode -v - | swiss base64 decode -v -
//	cat data.json | swiss json toYAML --value -
var stdinFlags = []string{"value", "url"}

// resolveStdin replaces any value flag set to "-" with the full contents of the
// command's stdin. It runs before every command via rootCmd.PersistentPreRunE.
func resolveStdin(cmd *cobra.Command) error {
	for _, name := range stdinFlags {
		f := cmd.Flags().Lookup(name)
		if f == nil || f.Value.String() != "-" {
			continue
		}
		b, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read stdin: %w", err)
		}
		if err := cmd.Flags().Set(name, string(b)); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return resolveStdin(cmd)
	}
}
