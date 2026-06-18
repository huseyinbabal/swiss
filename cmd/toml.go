package cmd

import (
	"os"
	"swiss/internal/toml"

	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
)

var tomlValue string

var tomlCmd = &cobra.Command{Use: "toml", Short: "toml related utility functions"}

var tomlToJSONCmd = &cobra.Command{
	Use:   "toJSON",
	Short: "convert toml to json",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := toml.ToJSON(tomlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "json", "terminal256", "monokai")
	},
}

var tomlToYAMLCmd = &cobra.Command{
	Use:   "toYAML",
	Short: "convert toml to yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := toml.ToYAML(tomlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "yaml", "terminal256", "monokai")
	},
}

var tomlFromJSONCmd = &cobra.Command{
	Use:   "fromJSON",
	Short: "convert json to toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := toml.FromJSON(tomlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "toml", "terminal256", "monokai")
	},
}

func init() {
	subCommands := []*cobra.Command{tomlToJSONCmd, tomlToYAMLCmd, tomlFromJSONCmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&tomlValue, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			panic("Please provide --value")
		}
		tomlCmd.AddCommand(c)
	}
	rootCmd.AddCommand(tomlCmd)
}
