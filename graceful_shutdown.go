package termtricks

import (
	"os"
	"os/signal"
	"syscall"
)

// GracefulShutdown and runs a custom function
func GracefulShutdown(fn func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		// Run user custom function
		fn()

		os.Exit(0)
	}()
}
