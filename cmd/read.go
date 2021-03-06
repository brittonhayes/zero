package cmd

import (
	"os"

	"github.com/brittonhayes/zero/internal/intel"
	"github.com/spf13/cobra"
)

// readCmd represents the go command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read from feed sources",
	Run: func(cmd *cobra.Command, args []string) {
		f := intel.New()
		results := make(chan intel.FeedItem)

		// Setup channel of feed
		// items
		f.Setup(results)

		// Fetch and print
		// the news
		f.Read(results, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
