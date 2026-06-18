package cmd

import (
	"fmt"
	"swiss/internal/lorem"

	"github.com/spf13/cobra"
)

var (
	loremWords      int
	loremSentences  int
	loremParagraphs int
)

var loremCmd = &cobra.Command{
	Use:   "lorem",
	Short: "generates lorem ipsum placeholder text",
	RunE: func(cmd *cobra.Command, args []string) error {
		if loremWords > 0 {
			fmt.Println(lorem.Words(loremWords))
		} else if loremSentences > 0 {
			fmt.Println(lorem.Sentences(loremSentences))
		} else {
			fmt.Println(lorem.Paragraphs(loremParagraphs))
		}
		return nil
	},
}

func init() {
	loremCmd.Flags().IntVar(&loremWords, "words", 0, "number of words to generate")
	loremCmd.Flags().IntVar(&loremSentences, "sentences", 0, "number of sentences to generate")
	loremCmd.Flags().IntVar(&loremParagraphs, "paragraphs", 1, "number of paragraphs to generate")
	rootCmd.AddCommand(loremCmd)
}
