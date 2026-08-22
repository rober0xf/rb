package browser

import (
	"os/exec"
)

func OpenBrowser(name string) error {
	urls := map[string]string{
		"chatgpt": "https://chatgpt.com",
		"github":  "https://github.com/rober0xf",
		"reddit":  "https://reddit.com",
		"claude":  "https://claude.ai/new",
		"youtube": "https://youtube.com",
		"mail":    "https://mail.google.com/mail/u/0/",
	}

	url, ok := urls[name]
	if !ok {
		return nil
	}

	return exec.Command("xdg-open", url).Start()
}