package password

import (
	"os"
	"os/exec"
)

func RunPass(args ...string) error {
	cmd := exec.Command("pass", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
