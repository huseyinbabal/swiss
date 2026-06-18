package cmd

import (
	"fmt"
	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
	"log"
	"os"
	"swiss/internal/json"
)

var jsonValue string

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json related utility functions",
}
var jsonBeautifyCmd = &cobra.Command{
	Use:   "beautify",
	Short: "Beautifies json",
	RunE: func(cmd *cobra.Command, args []string) error {
		json, err := json.Beautify(jsonValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, json, "json", "terminal256", "monokai")
	},
}

var jsonUglifyCmd = &cobra.Command{
	Use:   "uglify",
	Short: "Uglifies json",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.Uglify(jsonValue)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

var jsonToYamlCmd = &cobra.Command{
	Use:   "toYAML",
	Short: "Converts json to YAML",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.ToYAML(jsonValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "yaml", "terminal256", "monokai")
	},
}

var jsonToXmlCmd = &cobra.Command{
	Use:   "toXML",
	Short: "Converts json to XML",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.ToXML(jsonValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "xml", "terminal256", "monokai")
	},
}

var jsonToCsvCmd = &cobra.Command{
	Use:   "toCSV",
	Short: "Converts a json array of objects to CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.ToCSV(jsonValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "csv", "terminal256", "monokai")
	},
}

var jsonToTsvCmd = &cobra.Command{
	Use:   "toTSV",
	Short: "Converts a json array of objects to TSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.ToTSV(jsonValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var jsonToGoStructCmd = &cobra.Command{
	Use:   "toGoStruct",
	Short: "Generates a Go struct from json",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.ToGoStruct(jsonValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "go", "terminal256", "monokai")
	},
}

var jsonEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escapes a string into a json string literal",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.Escape(jsonValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var jsonUnescapeCmd = &cobra.Command{
	Use:   "unescape",
	Short: "Unescapes a json string literal",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := json.Unescape(jsonValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{
		jsonBeautifyCmd,
		jsonUglifyCmd,
		jsonToYamlCmd,
		jsonToXmlCmd,
		jsonToCsvCmd,
		jsonToTsvCmd,
		jsonToGoStructCmd,
		jsonEscapeCmd,
		jsonUnescapeCmd,
	}
	for _, c := range subCommands {
		c.Flags().StringVarP(&jsonValue, "value", "v", "", "JSON string")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide a valid json with --value parameter")
		}
		jsonCmd.AddCommand(c)
	}
	rootCmd.AddCommand(jsonCmd)
}
