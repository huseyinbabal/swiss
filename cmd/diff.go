package cmd

import (
	"fmt"
	"log"
	"swiss/internal/diff"

	"github.com/spf13/cobra"
)

var diffA string
var diffB string

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "compute a line-based diff between two inputs",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := diff.Lines(diffA, diffB)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	diffCmd.Flags().StringVar(&diffA, "a", "", "first input")
	if err := diffCmd.MarkFlagRequired("a"); err != nil {
		log.Fatalf("Please provide --a")
	}
	diffCmd.Flags().StringVar(&diffB, "b", "", "second input")
	if err := diffCmd.MarkFlagRequired("b"); err != nil {
		log.Fatalf("Please provide --b")
	}
	rootCmd.AddCommand(diffCmd)
}
