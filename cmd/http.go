package cmd

import (
	"github.com/brittonhayes/zero/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the http command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start zero http server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
