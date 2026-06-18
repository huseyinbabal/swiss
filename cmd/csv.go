package cmd

import (
	"fmt"
	"log"
	"os"
	"swiss/internal/csv"

	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
)

var csvValue string

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "csv related utility functions",
}

var csvToJsonCmd = &cobra.Command{
	Use:   "toJSON",
	Short: "Converts CSV into a json array of objects",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := csv.ToJSON(csvValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "json", "terminal256", "monokai")
	},
}

var csvToXmlCmd = &cobra.Command{
	Use:   "toXML",
	Short: "Converts CSV into XML",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := csv.ToXML(csvValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "xml", "terminal256", "monokai")
	},
}

var csvToYamlCmd = &cobra.Command{
	Use:   "toYAML",
	Short: "Converts CSV into YAML",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := csv.ToYAML(csvValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "yaml", "terminal256", "monokai")
	},
}

var csvEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escapes a value into a CSV field",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := csv.Escape(csvValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var csvUnescapeCmd = &cobra.Command{
	Use:   "unescape",
	Short: "Unescapes a CSV field into its literal value",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := csv.Unescape(csvValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{
		csvToJsonCmd,
		csvToXmlCmd,
		csvToYamlCmd,
		csvEscapeCmd,
		csvUnescapeCmd,
	}
	for _, c := range subCommands {
		c.Flags().StringVar(&csvValue, "value", "", "CSV string")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide a valid csv with --value parameter")
		}
		csvCmd.AddCommand(c)
	}
	rootCmd.AddCommand(csvCmd)
}
