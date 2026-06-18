package cmd

import (
	"fmt"
	"log"
	"swiss/internal/text"

	"github.com/spf13/cobra"
)

var textValue string

var textCmd = &cobra.Command{Use: "text", Short: "text related utility functions"}

var textCountCmd = &cobra.Command{
	Use:   "count",
	Short: "counts characters, bytes, words, lines and reading time",
	RunE: func(cmd *cobra.Command, args []string) error {
		res := text.Count(textValue)
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{textCountCmd}
	for _, c := range subCommands {
		c.Flags().StringVar(&textValue, "value", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		textCmd.AddCommand(c)
	}
	rootCmd.AddCommand(textCmd)
}
