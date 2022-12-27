package cmd

import (
	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
	"log"
	"os"
	"swiss/internal/xml"
)

var xmlValue string

var xmlCmd = &cobra.Command{
	Use:   "xml",
	Short: "xml related utility functions",
}
var xmlToJsonCmd = &cobra.Command{
	Use:   "toJSON",
	Short: "Converts XML into JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.ToJSON(xmlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, x, "json", "terminal256", "monokai")
	},
}

var xmlToYamlCmd = &cobra.Command{
	Use:   "toYAML",
	Short: "Converts XML into YAML",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.ToYAML(xmlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, x, "yaml", "terminal256", "monokai")
	},
}

var xmlToCsvCmd = &cobra.Command{
	Use:   "toCSV",
	Short: "Converts XML into CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.ToCSV(xmlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, x, "csv", "terminal256", "monokai")
	},
}

func init() {
	xmlToJsonCmd.Flags().StringVar(&xmlValue, "value", "", "XML string")
	err := xmlToJsonCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid xml with --value parameter")
	}
	xmlToYamlCmd.Flags().StringVar(&xmlValue, "value", "", "XML string")
	err = xmlToYamlCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid xml with --value parameter")
	}
	xmlToCsvCmd.Flags().StringVar(&xmlValue, "value", "", "XML string")
	err = xmlToCsvCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid xml with --value parameter")
	}
	xmlCmd.AddCommand(xmlToJsonCmd)
	xmlCmd.AddCommand(xmlToYamlCmd)
	xmlCmd.AddCommand(xmlToCsvCmd)
	rootCmd.AddCommand(xmlCmd)
}
