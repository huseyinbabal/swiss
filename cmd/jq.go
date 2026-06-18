package cmd

import (
	"os"
	"swiss/internal/jq"

	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
)

var jqPath string
var jqValue string

var jqCmd = &cobra.Command{
	Use:   "jq",
	Short: "query json with a dot/index path",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := jq.Query(jqPath, jqValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "json", "terminal256", "monokai")
	},
}

func init() {
	jqCmd.Flags().StringVar(&jqPath, "path", "", "path")
	jqCmd.Flags().StringVarP(&jqValue, "value", "v", "", "input")
	if err := jqCmd.MarkFlagRequired("path"); err != nil {
		panic("Please provide --path")
	}
	if err := jqCmd.MarkFlagRequired("value"); err != nil {
		panic("Please provide --value")
	}
	rootCmd.AddCommand(jqCmd)
}
