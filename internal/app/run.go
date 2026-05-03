package app

import (
	"io"
	"time"

	"github.com/fatih/color"
)

func RunWithSpinner(stderr io.Writer, message string, work func() error) error {
	frames := []string{"⣾", "⣷", "⣯", "⣟", "⣻", "⣽", "⣾", "⣷"}
	red := color.New(color.FgRed)
	purple := color.New(color.FgMagenta)

	done := make(chan error, 1)
	go func() {
		done <- work()
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case err := <-done:
			if err != nil {
				red.Fprintf(stderr, "Error: %v\n", err)
			}
			return err
		case <-ticker.C:
			purple.Fprintf(stderr, "\r%s %s", frames[i%len(frames)], message)
			i++
		}
	}
}