// Copyright 2019, Matteo Martellini
//
// Licensed under ISC, you can obtain a copy of the license
// inside the "LICENSE" file, on the root directory.
//
// +build ignore

package main

import (
	"fmt"
	"time"

	termtricks "github.com/mercxry/term-tricks"
)

func main() {
	// Hijacks SIGTERM (ctrl+c)
	termtricks.GracefulShutdown(safeExit)

	// Create an alternate buffer and hides the cursor
	termtricks.CreateAltBuffer()
	termtricks.HideCursor()

	// Actions on application exit
	defer func() {
		// Kills the alternate buffer and shows the cursor
		termtricks.KillAltBuffer()
		termtricks.ShowCursor()
	}()

	// Alternate buffer valid for 10 iterations (10s)
	for i := 5; i >= 1; i-- {
		// Reset cursor position to x=1, y=1
		termtricks.ResetCursorPos()

		fmt.Println("Hello World!")
		fmt.Println()

		// Handle the '1 second' typo
		if i == 1 {
			fmt.Printf("This message expires in... %d second\n", i)
		} else {
			fmt.Printf("This message expires in... %d seconds\n", i)
		}

		// Waits a second before going to the next loop iteration
		time.Sleep(time.Second)

		// Clears the terminal (in this case the alternate buffer screen)
		termtricks.Clear()
	}
}

// Actions to do for a clean exit
func safeExit() {
	// Kills the alternate buffer and shows the cursor
	termtricks.KillAltBuffer()
	termtricks.ShowCursor()
}
