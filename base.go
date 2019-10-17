package termtricks

import (
	"os"
	"os/exec"
)

// Clear ...
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
