package ai

import (
	"fmt"
	"time"
)

func spinner(done <-chan struct{}) {
	frames := []string{"Writing.", "Writing..", "Writing..."}
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r\033[K") // clear terminal
			return
		case <-ticker.C:
			fmt.Printf("\r\033[K%s", frames[i%len(frames)]) // clears the cursor
			i++
		}
	}
}
