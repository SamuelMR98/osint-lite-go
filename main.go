package main

import (
	"github.com/SamuelMR98/osint-lite-go/cmd"
	"github.com/SamuelMR98/osint-lite-go/internal"
	flag "github.com/spf13/pflag"
)

func main() {
	// Define flags with shorthand versions
	helpFlag := flag.BoolP("help", "h", false, "Show help message")
	socialFlag := flag.BoolP("social", "s", false, "Check social media platforms")
	techFlag := flag.BoolP("tech", "t", false, "Check tech platforms")
	jsonFlag := flag.BoolP("json", "j", false, "Output results in JSON format")
	flag.Parse()

	// Show help if the flag is set or if no username is provided
	if *helpFlag || flag.NArg() == 0 {
		cmd.Help()
		return
	}

	username := flag.Arg(0)

	// Determine which sites to check based on flags (run all if no specific category is selected)
	var selectedSites []internal.Site
	if *socialFlag {
		selectedSites = append(selectedSites, internal.GetSocialSites()...)
	}
	if *techFlag {
		selectedSites = append(selectedSites, internal.GetTechSites()...)
	}
	if !*socialFlag && !*techFlag {
		selectedSites = append(selectedSites, internal.GetSocialSites()...)
		selectedSites = append(selectedSites, internal.GetTechSites()...)
	}

	results := cmd.CheckSites(username, selectedSites)

	if *jsonFlag {
		cmd.PrintJSON(results)
	} else {
		cmd.PrintResults(results)
	}
}