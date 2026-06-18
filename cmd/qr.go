package cmd

import (
	"fmt"
	"log"
	"os"
	"swiss/internal/qr"

	"github.com/spf13/cobra"
)

var (
	qrValue  string
	qrOutput string
	qrSize   int
)

var qrCmd = &cobra.Command{
	Use:   "qr",
	Short: "Generates a QR code from a value",
	RunE: func(cmd *cobra.Command, args []string) error {
		if qrOutput != "" {
			png, err := qr.PNG(qrValue, qrSize)
			if err != nil {
				return err
			}
			if err := os.WriteFile(qrOutput, png, 0o644); err != nil {
				return err
			}
			fmt.Printf("QR code written to %s\n", qrOutput)
			return nil
		}
		out, err := qr.Terminal(qrValue)
		if err != nil {
			return err
		}
		fmt.Print(out)
		return nil
	},
}

func init() {
	qrCmd.Flags().StringVarP(&qrValue, "value", "v", "", "value to encode")
	qrCmd.Flags().StringVarP(&qrOutput, "output", "o", "", "write a PNG to this path instead of printing to the terminal")
	qrCmd.Flags().IntVar(&qrSize, "size", 256, "PNG size in pixels (used with --output)")
	if err := qrCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide a valid value with --value parameter")
	}
	rootCmd.AddCommand(qrCmd)
}
