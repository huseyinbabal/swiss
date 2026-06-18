package cmd

import (
	"fmt"
	"log"
	"swiss/internal/html"

	"github.com/spf13/cobra"
)

var htmlValue string

var htmlCmd = &cobra.Command{Use: "html", Short: "html related utility functions"}

var htmlEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "escape HTML special characters",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := html.Escape(htmlValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var htmlUnescapeCmd = &cobra.Command{
	Use:   "unescape",
	Short: "unescape HTML entities",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := html.Unescape(htmlValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var htmlStripCmd = &cobra.Command{
	Use:   "strip",
	Short: "strip HTML tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := html.Strip(htmlValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{htmlEscapeCmd, htmlUnescapeCmd, htmlStripCmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&htmlValue, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		htmlCmd.AddCommand(c)
	}
	rootCmd.AddCommand(htmlCmd)
}
