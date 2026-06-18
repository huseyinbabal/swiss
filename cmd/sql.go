package cmd

import (
	"fmt"
	"log"
	"os"
	"swiss/internal/sqlfmt"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/spf13/cobra"
)

var sqlValue string

var sqlCmd = &cobra.Command{Use: "sql", Short: "sql related utility functions"}

var sqlFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "format a SQL statement",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := sqlfmt.Format(sqlValue)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, res, "sql", "terminal256", "monokai")
	},
}

var sqlMinifyCmd = &cobra.Command{
	Use:   "minify",
	Short: "minify a SQL statement",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := sqlfmt.Minify(sqlValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{sqlFormatCmd, sqlMinifyCmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&sqlValue, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		sqlCmd.AddCommand(c)
	}
	rootCmd.AddCommand(sqlCmd)
}
