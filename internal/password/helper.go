package password

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunPass(args ...string) error {
	cmd := exec.Command("pass", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// prefixes the name with the current dir
func (m Model) passTitle(name string) string {
	rel := strings.TrimPrefix(m.currentDir, m.passwordsCommand.storeDir)
	rel = strings.TrimPrefix(rel, string(filepath.Separator))
	if rel == "" {
		return name
	}

	return filepath.Join(rel, name)
}

func (m Model) entryTitle(entry os.DirEntry) string {
	return m.passTitle(strings.TrimSuffix(entry.Name(), ".gpg"))
}

func (m Model) displayDir() string {
	rel := strings.TrimPrefix(m.currentDir, m.passwordsCommand.storeDir)
	rel = strings.TrimPrefix(rel, string(filepath.Separator))

	return rel
}
