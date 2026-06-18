package cmd

import (
	"fmt"
	"log"
	"swiss/internal/uuid"

	"github.com/spf13/cobra"
)

var uuidVersion int
var uuidValue string

var uuidCmd = &cobra.Command{Use: "uuid", Short: "uuid related utility functions"}

var uuidGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "generates a uuid (v4 or v7)",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			res string
			err error
		)
		switch uuidVersion {
		case 4:
			res, err = uuid.V4()
		case 7:
			res, err = uuid.V7()
		default:
			return fmt.Errorf("unsupported version %d (use 4 or 7)", uuidVersion)
		}
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var uuidValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validates a uuid",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := uuid.Validate(uuidValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	uuidGenCmd.Flags().IntVar(&uuidVersion, "version", 4, "uuid version (4 or 7)")

	uuidValidateCmd.Flags().StringVarP(&uuidValue, "value", "v", "", "uuid to validate")
	if err := uuidValidateCmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Please provide --value")
	}

	uuidCmd.AddCommand(uuidGenCmd)
	uuidCmd.AddCommand(uuidValidateCmd)
	rootCmd.AddCommand(uuidCmd)
}
