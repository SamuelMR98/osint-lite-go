package cmd

import (
	"fmt"

	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/SamuelMR98/osint-lite-go/internal/app"
	"github.com/spf13/cobra"
)


var checkOpts internal.CheckOptions

var checkCmd = &cobra.Command{
	Use: "check [username]",
	Short: "Check username across OSINT platforms",
	Long: `Check a username across various social media and tech platforms.
	By default, this command checks all supported platform categories.
	Use --social or --tech to limit the scope of the check.`,
	
	Example: ` osint-lite check jhonDoe
	osint-lite check jhonDoe --social
	osint-lite check jhonDoe --tech
	osint-lite check jhonDoe --json
	osint-lite check jhonDoe --json-save results.json
	osint-lite check jhonDoe -s -j -o results.json`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		selectedSites := app.BuildSelectedSites(checkOpts)

		var results []internal.Result

		checkWork := func() error {
			results = app.CheckSites(username, selectedSites)
			return nil
		}

		// Keep spinner off for JSON output (stdout remain clean)
		if checkOpts.NoSpinner || checkOpts.PrintJSON {
			if err := checkWork(); err != nil {
				return err
			}
		} else {
			if err := app.RunWithSpinner(cmd.ErrOrStderr(), "Checking sites", checkWork); err != nil {
				return err
			}
		}
		
		if checkOpts.PrintJSON {
			return app.PrintJSON(cmd.OutOrStdout(), results)
		}

		app.PrintResults(cmd.OutOrStdout(), results)

		if checkOpts.SaveJSON != "" {
			saveWork := func() error {
				return app.SaveJSON(results, checkOpts.SaveJSON)
			}
			if checkOpts.NoSpinner {
				if err := saveWork(); err != nil {
					return err
				}
			} else {
				if err := app.RunWithSpinner(cmd.ErrOrStderr(), "Saving results to JSON file...", saveWork); err != nil {
					return err
				}
			}
			fmt.Fprintf(cmd.ErrOrStderr(), "Saved results to %s\n", checkOpts.SaveJSON)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().BoolVarP(&checkOpts.Social, "social", "s", false, "Check social media platforms")
	checkCmd.Flags().BoolVarP(&checkOpts.Tech, "tech", "t", false, "Check tech platforms")
	checkCmd.Flags().BoolVarP(&checkOpts.PrintJSON, "json", "j", false, "Output results in JSON format")
	checkCmd.Flags().StringVarP(&checkOpts.SaveJSON, "json-save", "o", "", "Save results to a JSON file")
	checkCmd.Flags().BoolVarP(&checkOpts.NoSpinner, "no-spinner", "", false, "Disable spinner during checks")
}