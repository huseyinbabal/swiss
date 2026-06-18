package cmd

import (
	"fmt"
	"log"
	"swiss/internal/color"

	"github.com/spf13/cobra"
)

var colorValue string

var colorCmd = &cobra.Command{Use: "color", Short: "color related utility functions"}

var colorToRGBCmd = &cobra.Command{
	Use:   "toRGB",
	Short: "converts a hex color to rgb",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := color.ToRGB(colorValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var colorToHexCmd = &cobra.Command{
	Use:   "toHex",
	Short: "converts an rgb color to hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := color.ToHex(colorValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var colorToHSLCmd = &cobra.Command{
	Use:   "toHSL",
	Short: "converts a hex color to hsl",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := color.ToHSL(colorValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{colorToRGBCmd, colorToHexCmd, colorToHSLCmd}
	for _, c := range subCommands {
		c.Flags().StringVar(&colorValue, "value", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		colorCmd.AddCommand(c)
	}
	rootCmd.AddCommand(colorCmd)
}
