package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/fatih/color"
)

func PrintJSON(w io.Writer, results []internal.Result) error {
	blue := color.New(color.FgBlue)
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding results to JSON: %v", err)
	}

	blue.Fprintln(w, string(jsonData))
	return nil
}
func SaveJSON(results []internal.Result, filename string) error {
	blue := color.New(color.FgBlue)
	blue.Printf("Saving results to %s...\n", filename)
	red := color.New(color.FgRed)
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		red.Printf("Error encoding results to JSON: %v\n", err)
		return fmt.Errorf("error encoding results to JSON: %v", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		red.Printf("Error writing JSON to file: %v\n", err)
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	blue.Printf("Results saved to %s\n", filename)
	return nil
}
