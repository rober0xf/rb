package password

import (
	"os"
	"os/exec"
)

func runPass(args ...string) error {
	cmd := exec.Command("pass", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
