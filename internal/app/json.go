package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SamuelMR98/osint-lite-go/internal"
)

func PrintJSON(results []internal.Result) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding results to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func SaveJSON(filename string, results []internal.Result) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding results to JSON: %v\n", err)
		return
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error saving JSON to file: %v\n", err)
		return
	}
	fmt.Printf("Results saved to %s\n", filename)
}
