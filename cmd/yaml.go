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

var yamlToXmlCmd = &cobra.Command{
	Use:   "toXML",
	Short: "Converts yaml into XML",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := yaml.ToXML(yamlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "xml", "terminal256", "monokai")
	},
}

var yamlToCsvCmd = &cobra.Command{
	Use:   "toCSV",
	Short: "Converts a yaml array of objects into CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := yaml.ToCSV(yamlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "csv", "terminal256", "monokai")
	},
}

func init() {
	subCommands := []*cobra.Command{
		yamlBeautifyCmd,
		yamlToJsonCmd,
		yamlFormatCmd,
		yamlToXmlCmd,
		yamlToCsvCmd,
	}
	for _, c := range subCommands {
		c.Flags().StringVar(&yamlValue, "value", "", "YAML string")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide a valid yaml with --value parameter")
		}
		yamlCmd.AddCommand(c)
	}
	rootCmd.AddCommand(yamlCmd)
}
