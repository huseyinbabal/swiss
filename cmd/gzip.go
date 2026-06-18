package cmd

import (
	"fmt"
	"log"
	"swiss/internal/gzip"

	"github.com/spf13/cobra"
)

var gzipValue string

var gzipCmd = &cobra.Command{Use: "gzip", Short: "gzip related utility functions"}

var gzipCompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "gzip-compresses the input and base64-encodes it",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := gzip.Compress(gzipValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var gzipDecompressCmd = &cobra.Command{
	Use:   "decompress",
	Short: "base64-decodes then gunzips the input",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := gzip.Decompress(gzipValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{gzipCompressCmd, gzipDecompressCmd}
	for _, c := range subCommands {
		c.Flags().StringVar(&gzipValue, "value", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		gzipCmd.AddCommand(c)
	}
	rootCmd.AddCommand(gzipCmd)
}
