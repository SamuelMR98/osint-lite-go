package cmd

import (
	"net/http"

	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/fatih/color"
)

func PrintResults(results []internal.Result) {
	for _, result := range results {
		if result.Error != "" {
			color.Red("%s: Error checking %s - %s", result.Site, result.URL, result.Error)
		}
		if result.Found {
			color.Green("%s: Found at %s (Status Code: %d - %s)", result.Site, result.URL, result.StatusCode, http.StatusText(result.StatusCode))
		} else {
			color.Yellow("%s: Not found at %s (Status Code: %d - %s)", result.Site, result.URL, result.StatusCode, http.StatusText(result.StatusCode))
		}
	}
}