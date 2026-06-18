package cmd

import (
	"fmt"
	"log"
	"swiss/internal/timeconv"

	"github.com/spf13/cobra"
)

var timeValue string

var timeCmd = &cobra.Command{Use: "time", Short: "time related utility functions"}

var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "prints the current time in several formats",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := timeconv.Now()
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var timeToUnixCmd = &cobra.Command{
	Use:   "toUnix",
	Short: "converts a date string to unix seconds",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := timeconv.ToUnix(timeValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var timeFromUnixCmd = &cobra.Command{
	Use:   "fromUnix",
	Short: "converts unix seconds to an RFC3339 UTC date",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := timeconv.FromUnix(timeValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	valueCommands := []*cobra.Command{timeToUnixCmd, timeFromUnixCmd}
	for _, c := range valueCommands {
		c.Flags().StringVar(&timeValue, "value", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		timeCmd.AddCommand(c)
	}
	timeCmd.AddCommand(timeNowCmd)
	rootCmd.AddCommand(timeCmd)
}
