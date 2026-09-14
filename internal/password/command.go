package password

import (
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type PasswordManager struct{}

var initCmd = &cobra.Command{
	Use:   "init [gpg-id]",
	Short: "initialize the password store",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := PasswordManager{}
		return manager.InitStore(args[0])
	},
}

func (p *PasswordManager) InitStore(gpgID string) error {
	return RunPass("init", gpgID)
}

func (p *PasswordManager) ListPasswords() ([]Password, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	storePath := filepath.Join(home, ".password-store")
	var passwords []Password

	err = filepath.WalkDir(storePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".gpg" {
			return nil
		}

		relative, err := filepath.Rel(storePath, path)
		if err != nil {
			return err
		}

		title := strings.TrimSuffix(relative, ".gpg")
		passwords = append(passwords, Password{Title: title})

		return nil
	})

	return passwords, nil
}

func (p *PasswordManager) ShowPassword(title string) (Password, error) {
	output, err := exec.Command("pass", "show", title).Output()
	if err != nil {
		return Password{}, err
	}

	lines := strings.Split(strings.TrimSuffix(string(output), "\n"), "\n")

	password := Password{
		Title: title,
	}

	if len(lines) > 0 {
		password.Password = lines[0]
	}

	if len(lines) > 1 {
		password.Description = strings.Join(lines[1:], "\n")
	}

	return password, nil
}

func (p *PasswordManager) SavePassword(password Password) error {
	cmd := exec.Command("pass", "insert", "--multiline", password.Title)

	desc := password.Password
	if password.Description != "" {
		desc = "\n" + password.Description
	}

	cmd.Stdin = strings.NewReader(desc + "\n")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	return cmd.Run()
}

func (p *PasswordManager) DeletePassword(title string) error {
	cmd := exec.Command("pass", "rm", title)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	return cmd.Run()
}
