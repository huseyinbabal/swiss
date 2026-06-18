package cmd

import (
	"fmt"
	"log"
	"swiss/internal/base64"

	"github.com/spf13/cobra"
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
		fmt.Print(base64.Encode(b64Value))
		return nil
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes given base64 string",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := base64.Decode(b64Value)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	base64EncodeCmd.Flags().StringVarP(&b64Value, "value", "v", "", "string")
	err := base64EncodeCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid string with --value parameter")
	}
	base64DecodeCmd.Flags().StringVarP(&b64Value, "value", "v", "", "string")
	err = base64DecodeCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid string with --value parameter")
	}
	base64Cmd.AddCommand(base64EncodeCmd)
	base64Cmd.AddCommand(base64DecodeCmd)
	rootCmd.AddCommand(base64Cmd)
}
