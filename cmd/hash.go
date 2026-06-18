package cmd

import (
	"fmt"
	"log"
	"swiss/internal/hash"

	"github.com/spf13/cobra"
)

var hashValue string

var hashCmd = &cobra.Command{Use: "hash", Short: "hash related utility functions"}

var hashMD5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "computes the MD5 digest",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(hash.MD5(hashValue))
		return nil
	},
}

var hashSHA1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "computes the SHA-1 digest",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(hash.SHA1(hashValue))
		return nil
	},
}

var hashSHA256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "computes the SHA-256 digest",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(hash.SHA256(hashValue))
		return nil
	},
}

var hashSHA512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "computes the SHA-512 digest",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(hash.SHA512(hashValue))
		return nil
	},
}

var hashCRC32Cmd = &cobra.Command{
	Use:   "crc32",
	Short: "computes the IEEE CRC-32 checksum",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(hash.CRC32(hashValue))
		return nil
	},
}

func init() {
	subCommands := []*cobra.Command{hashMD5Cmd, hashSHA1Cmd, hashSHA256Cmd, hashSHA512Cmd, hashCRC32Cmd}
	for _, c := range subCommands {
		c.Flags().StringVarP(&hashValue, "value", "v", "", "input")
		if err := c.MarkFlagRequired("value"); err != nil {
			log.Fatalf("Please provide --value")
		}
		hashCmd.AddCommand(c)
	}
	rootCmd.AddCommand(hashCmd)
}
