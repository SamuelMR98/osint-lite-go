package cmd

// Help command implementation for osint-lite-go.
// This command provides users with information on how to use the tool, including available options and examples.

import (
	"github.com/fatih/color"
)

func Help() {
	color.Cyan("Usage: osint-lite <username>")
	color.Cyan("\nOptions:")
	color.Cyan("  -h, --help    Show this help message")
	color.Cyan("\nExamples:")
	color.Cyan("  osint-lite john_doe")
	color.Cyan("  osint-lite jane.smith")
}

