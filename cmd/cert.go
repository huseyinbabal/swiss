package cmd

import (
	"fmt"
	"log"
	"swiss/internal/cert"

	"github.com/spf13/cobra"
)

var certValue string

var certCmd = &cobra.Command{Use: "cert", Short: "cert related utility functions"}

var certInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "shows information about a PEM certificate",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := cert.Info(certValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{certInfoCmd}
	for _, c := range subCommands {
		c.Flags().StringVar(&certValue, "value", "", "PEM encoded certificate")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		certCmd.AddCommand(c)
	}
	rootCmd.AddCommand(certCmd)
}
