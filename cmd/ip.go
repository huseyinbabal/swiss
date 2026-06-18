package cmd

import (
	"fmt"
	"swiss/internal/ip"

	"github.com/spf13/cobra"
)

var ipValue string

var ipCmd = &cobra.Command{Use: "ip", Short: "ip related utility functions"}

var ipCidrCmd = &cobra.Command{
	Use:   "cidr",
	Short: "report on a CIDR block",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := ip.CIDR(ipValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var ipToIntCmd = &cobra.Command{
	Use:   "toInt",
	Short: "convert IPv4 to integer",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := ip.ToInt(ipValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var ipFromIntCmd = &cobra.Command{
	Use:   "fromInt",
	Short: "convert integer to IPv4",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := ip.FromInt(ipValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{ipCidrCmd, ipToIntCmd, ipFromIntCmd}
	for _, c := range subCommands {
		c.Flags().StringVar(&ipValue, "value", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			panic("Please provide --value")
		}
		ipCmd.AddCommand(c)
	}
	rootCmd.AddCommand(ipCmd)
}
