package cmd

import (
	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
	"os"
	"swiss/internal/password"
)

var includeUpper bool
var includeNum bool
var includeSymbol bool
var passwordLength int

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "password generator",
}
var passwordGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates password",
	RunE: func(cmd *cobra.Command, args []string) error {
		pass, err := password.Generate(passwordLength, includeUpper, includeNum, includeSymbol)
		if err != nil {
			return err
		}
		return quick.Highlight(os.Stdout, pass, "text", "terminal256", "monokai")
	},
}

func init() {
	passwordGenerateCmd.Flags().BoolVar(&includeUpper, "include-upper", false, "Include uppercase characters")
	passwordGenerateCmd.Flags().BoolVar(&includeNum, "include-numeric", false, "Include numeric characters")
	passwordGenerateCmd.Flags().BoolVar(&includeSymbol, "include-symbol", false, "Include special characters")
	passwordGenerateCmd.Flags().IntVar(&passwordLength, "length", 16, "Password length")

	passwordCmd.AddCommand(passwordGenerateCmd)
	rootCmd.AddCommand(passwordCmd)
}
