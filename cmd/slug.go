package cmd

import (
	"fmt"
	"log"
	"swiss/internal/slug"

	"github.com/spf13/cobra"
)

var slugValue string

var slugCmd = &cobra.Command{
	Use:   "slug",
	Short: "generates a URL-friendly slug from the input",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(slug.Make(slugValue))
		return nil
	},
}

func init() {
	slugCmd.Flags().StringVar(&slugValue, "value", "", "input")
	if err := slugCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide --value")
	}
	rootCmd.AddCommand(slugCmd)
}
