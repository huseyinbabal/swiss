package cmd

import (
	"fmt"
	"log"
	"swiss/internal/jwt"

	"github.com/spf13/cobra"
)

var jwtValue string
var jwtSecret string

var jwtCmd = &cobra.Command{Use: "jwt", Short: "jwt related utility functions"}

var jwtDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decodes a JWT token",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := jwt.Decode(jwtValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var jwtVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verifies a JWT HS256 signature",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := jwt.Verify(jwtValue, jwtSecret)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{jwtDecodeCmd, jwtVerifyCmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&jwtValue, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		jwtCmd.AddCommand(c)
	}
	jwtVerifyCmd.Flags().StringVar(&jwtSecret, "secret", "", "secret")
	if err := jwtVerifyCmd.MarkFlagRequired("secret"); err != nil {
		log.Fatalf("Please provide --secret")
	}
	rootCmd.AddCommand(jwtCmd)
}
