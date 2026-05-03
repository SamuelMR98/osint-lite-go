package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const appVersion = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "osint-lite-go",
	Short: "A lightweight OSINT tool for checking username availability across various platforms.",
	Long: `osint-lite-go is a command-line tool designed to help users quickly check the availability of usernames across a wide range of social media and tech platforms. 
It provides a simple interface for OSINT investigations, allowing users to gather information about potential online identities.`,
	Version: appVersion,

	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() error {
	red := color.New(color.FgRed)
	if err := rootCmd.Execute(); err != nil {
		red.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	return nil
}

func init() {
	rootCmd.SetVersionTemplate("osint-lite-go version {{.Version}}\n")
}