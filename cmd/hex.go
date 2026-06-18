package cmd

import (
	"fmt"
	"log"
	"swiss/internal/hex"

	"github.com/spf13/cobra"
)

var hexValue string

var hexCmd = &cobra.Command{Use: "hex", Short: "hex related utility functions"}

var hexEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "hex-encodes the input",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := hex.Encode(hexValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var hexDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "hex-decodes the input",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := hex.Decode(hexValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{hexEncodeCmd, hexDecodeCmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&hexValue, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		hexCmd.AddCommand(c)
	}
	rootCmd.AddCommand(hexCmd)
}
