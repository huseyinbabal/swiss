package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"swiss/internal/url"
)

var urlString string

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "url related functions",
}
var urlEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encodes url",
	RunE: func(cmd *cobra.Command, args []string) error {
		encodedUrl, err := url.Encode(urlString)
		if err != nil {
			return err
		}
		fmt.Print(encodedUrl)
		return nil
	},
}

var urlDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes url",
	RunE: func(cmd *cobra.Command, args []string) error {
		decodedUrl, err := url.Decode(urlString)
		if err != nil {
			return err
		}
		fmt.Print(decodedUrl)
		return nil
	},
}

func init() {
	urlEncodeCmd.Flags().StringVar(&urlString, "url", "", "URL")
	err := urlEncodeCmd.MarkFlagRequired("url")
	if err != nil {
		log.Fatalf("Please provide a valid url --url parameter")
	}
	urlDecodeCmd.Flags().StringVar(&urlString, "url", "", "URL")
	err = urlDecodeCmd.MarkFlagRequired("url")
	if err != nil {
		log.Fatalf("Please provide a valid url --url parameter")
	}
	urlCmd.AddCommand(urlEncodeCmd)
	urlCmd.AddCommand(urlDecodeCmd)
	rootCmd.AddCommand(urlCmd)
}
