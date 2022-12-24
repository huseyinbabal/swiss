package cmd

import (
	"fmt"
	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
	"log"
	"os"
	"swiss/internal/yaml"
)

var yamlValue string

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "yaml related utility functions",
}
var yamlBeautifyCmd = &cobra.Command{
	Use:   "beautify",
	Short: "Beautifies yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		json, err := yaml.Beautify(yamlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, json, "yaml", "terminal256", "monokai")
	},
}

var yamlFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "Formats yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		y, err := yaml.Format(yamlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, y, "yaml", "terminal256", "monokai")
	},
}
var yamlToJsonCmd = &cobra.Command{
	Use:   "toJson",
	Short: "Converts yaml into json",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := yaml.ToJson(yamlValue)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	yamlBeautifyCmd.Flags().StringVar(&yamlValue, "value", "", "YAML string")
	err := yamlBeautifyCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid yaml with --value parameter")
	}
	yamlToJsonCmd.Flags().StringVar(&yamlValue, "value", "", "YAML string")
	err = yamlToJsonCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid yaml with --value parameter")
	}
	yamlFormatCmd.Flags().StringVar(&yamlValue, "value", "", "YAML string")
	err = yamlFormatCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid yaml with --value parameter")
	}
	yamlCmd.AddCommand(yamlBeautifyCmd)
	yamlCmd.AddCommand(yamlToJsonCmd)
	yamlCmd.AddCommand(yamlFormatCmd)
	rootCmd.AddCommand(yamlCmd)
}
