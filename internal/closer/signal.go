package closer

import (
	"context"
	"os"
	"os/signal"
)

// CloseOnSignal wait one of signals and calls Close.
func CloseOnSignal(ctx context.Context, signals ...os.Signal) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, signals...)

	for {
		select {
		case <-ctx.Done():
			return Close()
		case <-done:
			return Close()
		}
	}
}
