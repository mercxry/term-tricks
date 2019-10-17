package termtricks

import "fmt"

// MoveCursor ...
func MoveCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

// ResetCursorPos ...
func ResetCursorPos() {
	fmt.Print("\033[1;1H")
}

// HideCursor ...
func HideCursor() {
	fmt.Print("\033[?25l")
}

// ShowCursor ...
func ShowCursor() {
	fmt.Print("\033[?25h")
}
