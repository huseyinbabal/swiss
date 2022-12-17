package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "swiss",
	Short: "swiss is a command that contains utility functions like encode, decode, etc...",
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
