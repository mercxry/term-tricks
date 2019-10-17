# Term(inal) Tricks
[![asciicast](https://asciinema.org/a/nCVomtIl8QUTbn5KXUp0Zu6nH.svg)](https://asciinema.org/a/nCVomtIl8QUTbn5KXUp0Zu6nH)

Cool library for cool terminal stuff, it's a mixed library, more towards a personal use, but don't fear pushing merge requests!

**Supports only Linux and MacOS**

## Table of Content
- [Installation](#installation)
- [Features](#features)
- [Example](#example)

## Installation
```sh
go get -u github.com/mercxry/term-tricks
```

## Features
- **Create** and **Kill** an Alternate Buffer
- **Hide** and **Show** the Terminal Cursor
- **Some** Commands Bindings for Golang (**I.E:** *clear*)
- **More** *coming soon!*

## Example
```go
// Copyright 2019, Matteo Martellini
//
// Licensed under ISC, you can obtain a copy of the license
// inside the "LICENSE" file, on the root directory.

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
```
- [More examples in the "examples" directory](https://github.com/mercxry/term-tricks/tree/master/examples)

## License
[ISC](https://github.com/mercxry/term-tricks/blob/master/LICENSE)