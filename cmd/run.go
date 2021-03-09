package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/brittonhayes/zero/templates"
	"github.com/brittonhayes/zero/zero"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the go command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Fetch all matching zero day feed results",
	Run: func(cmd *cobra.Command, args []string) {
		z := zero.Setup()
		matches, err := z.ReadRSS().Inspect()
		if err != nil {
			log.Error(err)
			return
		}

		// Content results in preferred format
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
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("template", "t", "", "render output with custom template")
	runCmd.Flags().BoolP("json", "j", false, "render output as raw JSON")
	runCmd.Flags().Bool("check", false, "Exit with status 1 if results are found")

	_ = viper.BindPFlag("check", runCmd.Flags().Lookup("check"))
	_ = viper.BindPFlag("json", runCmd.Flags().Lookup("json"))
}
