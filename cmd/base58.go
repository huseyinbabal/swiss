package cmd

import (
	"fmt"
	"log"
	"swiss/internal/base58"

	"github.com/spf13/cobra"
)

var base58Value string

var base58Cmd = &cobra.Command{Use: "base58", Short: "base58 related utility functions"}

var base58EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encodes input to base58",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := base58.Encode(base58Value)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var base58DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decodes base58 input",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := base58.Decode(base58Value)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{base58EncodeCmd, base58DecodeCmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&base58Value, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		base58Cmd.AddCommand(c)
	}
	rootCmd.AddCommand(base58Cmd)
}
