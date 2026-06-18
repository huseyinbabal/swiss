package cmd

import (
	"fmt"
	"log"
	"swiss/internal/bcrypt"

	"github.com/spf13/cobra"
)

var bcryptValue string
var bcryptPassword string

var bcryptCmd = &cobra.Command{Use: "bcrypt", Short: "bcrypt related utility functions"}

var bcryptHashCmd = &cobra.Command{
	Use:   "hash",
	Short: "hashes a password with bcrypt",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := bcrypt.Hash(bcryptValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var bcryptVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verifies a password against a bcrypt hash",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := bcrypt.Verify(bcryptValue, bcryptPassword)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	bcryptHashCmd.Flags().StringVar(&bcryptValue, "value", "", "password to hash")
	if err := bcryptHashCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide --value")
	}

	bcryptVerifyCmd.Flags().StringVar(&bcryptValue, "value", "", "bcrypt hash")
	bcryptVerifyCmd.Flags().StringVar(&bcryptPassword, "password", "", "password to verify")
	if err := bcryptVerifyCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide --value")
	}
	if err := bcryptVerifyCmd.MarkFlagRequired("password"); err != nil {
		log.Fatalf("Please provide --password")
	}

	bcryptCmd.AddCommand(bcryptHashCmd)
	bcryptCmd.AddCommand(bcryptVerifyCmd)
	rootCmd.AddCommand(bcryptCmd)
}
