package cmd

import (
	"fmt"

	"github.com/brittonhayes/zero/internal/intel"
	"github.com/spf13/cobra"
)

// readCmd represents the go command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "ReadRSS from feed sources",
	Run: func(cmd *cobra.Command, args []string) {
		config := intel.New()
		rss := config.ReadRSS()
		matches, err := rss.FindMatches()
		if err != nil {
			panic(err)
		}

		for _, m := range matches {
			fmt.Printf("%+v\n", m.Item.Title)
		}
		//
		// switch cmd.Flag("template").Value.String() != "" {
		// case true:
		// 	// fetch and print the news
		// 	//  with a custom template
		// 	f.ReadTemplate(results, os.Stdout, cmd.Flag("template").Value.String())
		// default:
		// 	// fetch and print
		// 	// the news
		// 	f.ReadRSS(results, os.Stdout)
		// }
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringP("template", "t", "", "render output with custom template")
}
