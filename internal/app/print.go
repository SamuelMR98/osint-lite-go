package app

import (
	"io"

	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/SamuelMR98/osint-lite-go/utils"
	"github.com/fatih/color"
)

func PrintResults(w io.Writer, results []internal.Result) {
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	for _, result := range results {
		if result.Found {
			green.Fprintf(w, "✓ %s: %s (Status: %d - %s)\n", result.Site, result.URL, result.StatusCode, utils.GetStatusText(result.StatusCode))
		} else if result.Error != "" {
			red.Fprintf(w, "✗ %s: %s (Error: %s)\n", result.Site, result.URL, result.Error)
		} else {
			yellow.Fprintf(w, "⚠ %s: %s (Status: %d - %s)\n", result.Site, result.URL, result.StatusCode, utils.GetStatusText(result.StatusCode))
		}
	}
}
