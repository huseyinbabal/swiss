package cmd

import (
	"fmt"
	"log"
	"os"
	"swiss/internal/xml"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/spf13/cobra"
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

var xmlBeautifyCmd = &cobra.Command{
	Use:   "beautify",
	Short: "Beautifies XML with indentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.Beautify(xmlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, x, "xml", "terminal256", "monokai")
	},
}

var xmlUglifyCmd = &cobra.Command{
	Use:   "uglify",
	Short: "Removes insignificant whitespace from XML",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.Uglify(xmlValue)
		if err != nil {
			return err
		}
		fmt.Println(x)
		return nil
	},
}

var xmlEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escapes XML-significant characters",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.Escape(xmlValue)
		if err != nil {
			return err
		}
		fmt.Println(x)
		return nil
	},
}

var xmlUnescapeCmd = &cobra.Command{
	Use:   "unescape",
	Short: "Unescapes XML entity references",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := xml.Unescape(xmlValue)
		if err != nil {
			return err
		}
		fmt.Println(x)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{
		xmlToJsonCmd,
		xmlToYamlCmd,
		xmlToCsvCmd,
		xmlBeautifyCmd,
		xmlUglifyCmd,
		xmlEscapeCmd,
		xmlUnescapeCmd,
	}
	for _, c := range subCommands {
		c.Flags().StringVarP(&xmlValue, "value", "v", "", "XML string")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide a valid xml with --value parameter")
		}
		xmlCmd.AddCommand(c)
	}
	rootCmd.AddCommand(xmlCmd)
}
