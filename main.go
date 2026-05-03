package main

import (
	"fmt"
	"io"
	"os"

	"github.com/SamuelMR98/osint-lite-go/cmd"
	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

const (
	appName = "osint-lite-go"
	version = "0.1.0"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		fmt.Fprintf(os.Stderr, "Run '%s --help' for usage. \n", appName)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	cyan := color.New(color.FgCyan)
	green := color.New(color.FgGreen)

	cfg, fs, err := parseFlags(args, stdout)
	if err != nil {
		return err
	}

	if cfg.ShowHelp {
		fs.Usage()
		return nil
	}

	if cfg.Version {
		cyan.Fprintf(stdout, "%s version %s\n", appName, version)
		return nil
	}

	if cfg.Username == "" {
		fs.Usage()
		return fmt.Errorf("username is required")
	}

	selectedSites := cmd.BuildSelectedSites(cfg)

	var results []internal.Result

	checkWork := func() error {
		results = cmd.CheckSites(cfg.Username, selectedSites)
		return nil
	}

	if !cfg.NoSpinner || cfg.PrintJSON {
		// Avoid spinner if JSON output is requested
		if err := checkWork(); err != nil {
			return err
		}
	} else {
		// Show spinner for non-JSON output
		if err := cmd.RunWithSpinner(stderr, "Checking username availability...", checkWork); err != nil {
			return err
		}
	}

	if cfg.PrintJSON {
		cmd.PrintJSON(results)
	} else {
		cmd.PrintResults(results)
	}

	if cfg.SaveJSON != "" {
		saveWork := func() error {
			cmd.SaveJSON(cfg.SaveJSON, results)
			return nil
		}
		if !cfg.NoSpinner {
			if err := cmd.RunWithSpinner(stderr, "Saving results to JSON...", saveWork); err != nil {
				return err
			}
		} else {
			if err := saveWork(); err != nil {
				return err
			}
		}
		green.Fprintf(stderr, "Results saved to %s\n", cfg.SaveJSON)
	}

	return nil
}

func parseFlags(args []string, stdout io.Writer) (*internal.Config, *flag.FlagSet, error) {
	var cfg internal.Config
	
	fs := flag.NewFlagSet(appName, flag.ContinueOnError)
	fs.SetOutput(stdout)

	fs.BoolVarP(&cfg.ShowHelp, "help", "h", false, "Show help message")
	fs.BoolVarP(&cfg.Version, "version", "v", false, "Show version information")

	fs.BoolVarP(&cfg.Social, "social", "s", false, "Check social media platforms")
	fs.BoolVarP(&cfg.Tech, "tech", "t", false, "Check tech platforms")

	fs.BoolVarP(&cfg.PrintJSON, "json", "j", false, "Output results in JSON format")
	fs.StringVarP(&cfg.SaveJSON, "json-save", "", "", "Save results to a JSON file")

	fs.BoolVarP(&cfg.NoSpinner, "no-spinner", "", false, "Disable spinner during checks and saving")

	fs.Usage = func() {
		cmd.Help()
	}

	if err := fs.Parse(args); err != nil {
		return &cfg, fs, err
	}

	if fs.NArg() > 0 {
		cfg.Username = fs.Arg(0)
	}

	return &cfg, fs, nil
}
