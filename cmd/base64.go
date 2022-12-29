package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"swiss/internal/base64"
)

var b64Value string

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64 related utility functions",
}
var base64EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encodes given string with base64",
	RunE: func(cmd *cobra.Command, args []string) error {
		if b64Value == "-" {
			inputReader := cmd.InOrStdin()
			b, err := io.ReadAll(inputReader)
			if err != nil {
				return err
			}
			b64Value = string(b)
		}
		fmt.Print(base64.Encode(b64Value))
		return nil
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes given base64 string",
	RunE: func(cmd *cobra.Command, args []string) error {
		if b64Value == "-" {
			inputReader := cmd.InOrStdin()
			b, err := io.ReadAll(inputReader)
			if err != nil {
				return err
			}
			b64Value = string(b)
		}
		res, err := base64.Decode(b64Value)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	base64EncodeCmd.Flags().StringVar(&b64Value, "value", "", "string")
	err := base64EncodeCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid string with --value parameter")
	}
	base64DecodeCmd.Flags().StringVar(&b64Value, "value", "", "string")
	err = base64DecodeCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid string with --value parameter")
	}
	base64Cmd.AddCommand(base64EncodeCmd)
	base64Cmd.AddCommand(base64DecodeCmd)
	rootCmd.AddCommand(base64Cmd)
}
