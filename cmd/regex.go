package cmd

import (
	"fmt"
	"log"
	"swiss/internal/regex"

	"github.com/spf13/cobra"
)

var regexPattern string
var regexValue string
var regexReplacement string

var regexCmd = &cobra.Command{Use: "regex", Short: "regex related utility functions"}

var regexMatchCmd = &cobra.Command{
	Use:   "match",
	Short: "find matches and capture groups",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := regex.Match(regexPattern, regexValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var regexReplaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "replace matches with a replacement",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := regex.Replace(regexPattern, regexValue, regexReplacement)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	regexMatchCmd.Flags().StringVar(&regexPattern, "pattern", "", "regex pattern")
	if err := regexMatchCmd.MarkFlagRequired("pattern"); err != nil {
		log.Fatalf("Please provide --pattern")
	}
	regexMatchCmd.Flags().StringVarP(&regexValue, "value", "v", "", "input")
	if err := regexMatchCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide --value")
	}

	regexReplaceCmd.Flags().StringVar(&regexPattern, "pattern", "", "regex pattern")
	if err := regexReplaceCmd.MarkFlagRequired("pattern"); err != nil {
		log.Fatalf("Please provide --pattern")
	}
	regexReplaceCmd.Flags().StringVarP(&regexValue, "value", "v", "", "input")
	if err := regexReplaceCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide --value")
	}
	regexReplaceCmd.Flags().StringVar(&regexReplacement, "replacement", "", "replacement string")
	if err := regexReplaceCmd.MarkFlagRequired("replacement"); err != nil {
		log.Fatalf("Please provide --replacement")
	}

	regexCmd.AddCommand(regexMatchCmd)
	regexCmd.AddCommand(regexReplaceCmd)
	rootCmd.AddCommand(regexCmd)
}
