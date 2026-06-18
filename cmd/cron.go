package cmd

import (
	"fmt"
	"swiss/internal/cron"

	"github.com/spf13/cobra"
)

var cronValue string
var cronCount int

var cronCmd = &cobra.Command{Use: "cron", Short: "cron related utility functions"}

var cronExplainCmd = &cobra.Command{
	Use:   "explain",
	Short: "explain a cron expression",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := cron.Explain(cronValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var cronNextCmd = &cobra.Command{
	Use:   "next",
	Short: "show the next matching times for a cron expression",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := cron.Next(cronValue, cronCount)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	cronExplainCmd.Flags().StringVar(&cronValue, "value", "", "input")
	if err := cronExplainCmd.MarkFlagRequired("value"); err != nil {
		panic("Please provide --value")
	}
	cronCmd.AddCommand(cronExplainCmd)

	cronNextCmd.Flags().StringVar(&cronValue, "value", "", "input")
	cronNextCmd.Flags().IntVar(&cronCount, "count", 5, "number of times to show")
	if err := cronNextCmd.MarkFlagRequired("value"); err != nil {
		panic("Please provide --value")
	}
	cronCmd.AddCommand(cronNextCmd)

	rootCmd.AddCommand(cronCmd)
}
