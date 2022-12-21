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

func init() {
	jsonBeautifyCmd.Flags().StringVar(&jsonValue, "value", "", "JSON string")
	err := jsonBeautifyCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid json with --value parameter")
	}
	jsonUglifyCmd.Flags().StringVar(&jsonValue, "value", "", "JSON string")
	err = jsonUglifyCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid json with --value parameter")
	}
	jsonToYamlCmd.Flags().StringVar(&jsonValue, "value", "", "JSON string")
	err = jsonToYamlCmd.MarkFlagRequired("value")
	if err != nil {
		log.Fatalf("Please provide a valid json with --value parameter")
	}
	jsonCmd.AddCommand(jsonBeautifyCmd)
	jsonCmd.AddCommand(jsonUglifyCmd)
	jsonCmd.AddCommand(jsonToYamlCmd)
	rootCmd.AddCommand(jsonCmd)
}
