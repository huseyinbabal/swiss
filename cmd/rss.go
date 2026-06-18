package cmd

import (
	"log"
	"os"
	"swiss/internal/rss"

	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
)

var rssValue string

var rssCmd = &cobra.Command{
	Use:   "rss",
	Short: "rss related utility functions",
}

var rssToJsonCmd = &cobra.Command{
	Use:   "toJSON",
	Short: "Converts an RSS feed into JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := rss.ToJSON(rssValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "json", "terminal256", "monokai")
	},
}

func init() {
	rssToJsonCmd.Flags().StringVarP(&rssValue, "value", "v", "", "RSS string")
	if err := rssToJsonCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide a valid rss with --value parameter")
	}
	rssCmd.AddCommand(rssToJsonCmd)
	rootCmd.AddCommand(rssCmd)
}
