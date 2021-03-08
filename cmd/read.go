package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/brittonhayes/zero/internal/zero"
	"github.com/brittonhayes/zero/templates"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// readCmd represents the go command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "ReadRSS from feed sources",
	Run: func(cmd *cobra.Command, args []string) {
		z := zero.Setup()
		matches, err := z.ReadRSS().Inspect()
		if err != nil {
			log.Error(err)
			return
		}

		// Render results in preferred format
		output(viper.GetBool("json"), matches)

		// Run in CI mode if enabled
		exitCheck(viper.GetBool("check"), len(matches))
	},
}

func output(enabled bool, m zero.Matches) {
	if enabled {
		b, err := m.MarshalJSON()
		if err != nil {
			log.Error(err)
		}
		fmt.Println(string(b))
		return
	}

	log.WithFields(log.Fields{
		"STATUS": "Complete",
	}).Debug()

	tpl, err := template.New("zero").Parse(templates.DefaultTemplate)
	if err != nil {
		log.Error(err)
		return
	}

	err = tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Error(err)
		return
	}

}

func exitCheck(enabled bool, count int) {
	if enabled && count > 0 {
		log.WithFields(log.Fields{
			"STATUS":  "failure",
			"MATCHES": count,
			"MESSAGE": "failed continuous integration check",
		}).Error()
		log.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringP("template", "t", "", "render output with custom template")
	readCmd.Flags().BoolP("json", "j", false, "render output as raw JSON")
	readCmd.Flags().Bool("check", false, "Exit with status 1 if results are found")

	_ = viper.BindPFlag("check", readCmd.Flags().Lookup("check"))
	_ = viper.BindPFlag("json", readCmd.Flags().Lookup("json"))
}
