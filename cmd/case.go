package cmd

import (
	"fmt"
	"log"
	"swiss/internal/textcase"

	"github.com/spf13/cobra"
)

var caseValue string

var caseCmd = &cobra.Command{Use: "case", Short: "case conversion utility functions"}

var caseCamelCmd = &cobra.Command{
	Use:   "camel",
	Short: "converts to camelCase",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToCamel(caseValue))
		return nil
	},
}

var casePascalCmd = &cobra.Command{
	Use:   "pascal",
	Short: "converts to PascalCase",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToPascal(caseValue))
		return nil
	},
}

var caseSnakeCmd = &cobra.Command{
	Use:   "snake",
	Short: "converts to snake_case",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToSnake(caseValue))
		return nil
	},
}

var caseKebabCmd = &cobra.Command{
	Use:   "kebab",
	Short: "converts to kebab-case",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToKebab(caseValue))
		return nil
	},
}

var caseTitleCmd = &cobra.Command{
	Use:   "title",
	Short: "converts to Title Case",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToTitle(caseValue))
		return nil
	},
}

var caseUpperCmd = &cobra.Command{
	Use:   "upper",
	Short: "converts to UPPERCASE",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToUpper(caseValue))
		return nil
	},
}

var caseLowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "converts to lowercase",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(textcase.ToLower(caseValue))
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{
		caseCamelCmd,
		casePascalCmd,
		caseSnakeCmd,
		caseKebabCmd,
		caseTitleCmd,
		caseUpperCmd,
		caseLowerCmd,
	}
	for _, c := range subCommands {
		c.Flags().StringVar(&caseValue, "value", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		caseCmd.AddCommand(c)
	}
	rootCmd.AddCommand(caseCmd)
}
