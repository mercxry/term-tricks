package termtricks

import "fmt"

// CreateAltBuffer ...
func CreateAltBuffer() {
	fmt.Print("\033[?1049h\033[H")
}

// CloseAltBuffer ...
func CloseAltBuffer() {
	fmt.Print("\033[?1049l")
}
