package app

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func RunWithSpinner(stderr interface{ Write([]byte) (int, error) }, message string, work func() error) error {
	purple := color.New(color.FgMagenta)
	frames := []string{"⣾", "⣷", "⣯", "⣟", "⣻", "⣽", "⣾", "⣷"}
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
			fmt.Fprintf(stderr, "\r%s %s\n", color.GreenString("✓"), message)
			return err
		case <-ticker.C:
			purple.Fprintf(stderr, "\r%s %s", frames[i%len(frames)], message)
			i++
		}
	}
}
