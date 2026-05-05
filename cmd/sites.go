package cmd

import (
	"github.com/SamuelMR98/osint-lite-go/internal/app"
	"github.com/spf13/cobra"
)

var sitesCmd = &cobra.Command{
	Use: "sites",
	Short: "List all supported OSINT platforms",
	Long: `List all supported OSINT platforms that osint-lite can check.
	This command is useful to see the full list of platforms and their categories.`,

	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		platforms := app.GetSupportedSites()
		for category, sites := range platforms {
			cmd.Printf("%s:\n", category)
			for _, site := range sites {
				cmd.Printf("  - %s\n", site)
			}
			cmd.Println()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sitesCmd)
}
