package main

import (
	"os"

	"github.com/SamuelMR98/osint-lite-go/cmd"
)


func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

