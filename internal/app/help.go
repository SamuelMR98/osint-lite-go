package app

// Help command implementation for osint-lite-go.
// This command provides users with information on how to use the tool, including available options and examples.

import (
	"github.com/fatih/color"
)

func Help() {
	color.Cyan("osint-lite-go - A simple OSINT tool for checking username availability across various platforms.\n")
	color.Green("Usage:")
	color.White("  osint-lite-go [options] <username>\n")
	color.Green("Options:")
	color.White("  -h, --help       Show help message")
	color.White("  -s, --social     Check social media platforms")
	color.White("  -t, --tech       Check tech platforms")
	color.White("  -j, --json       Output results in JSON format\n")
	color.Green("Examples:")
	color.White("  osint-lite-go johndoe")
	color.White("  osint-lite-go -s johndoe")
	color.White("  osint-lite-go -t johndoe")
	color.White("  osint-lite-go -s -t johndoe")
	color.White("  osint-lite-go -j johndoe\n")
	color.Cyan("For more information, visit the GitHub repository: https://github.com/SamuelMR98/osint-lite-go")
}
