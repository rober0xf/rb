package ai

import (
	"context"
	"fmt"

	appConfig "rb/config"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

const behavior = `
You are rb, a personal AI assistant for the terminal.
Be concise and practical. When helping with software engineering, debugging, logs, and code reviews.
`

func CallModel(ctx context.Context, cfg *appConfig.Config, prompt string) error {
	done := make(chan struct{})
	go spinner(done)

	s := openrouter.New(
		openrouter.WithSecurity(cfg.APIKey),
	)

	res, err := s.Chat.Send(ctx, components.ChatRequest{
		Model: new("deepseek/deepseek-v4-flash-0731:free"),
		Messages: []components.ChatMessages{
			// behavior
			components.CreateChatMessagesSystem(
				components.ChatSystemMessage{
					Content: components.CreateChatSystemMessageContentStr(behavior),
					Role:    components.ChatSystemMessageRoleSystem,
				},
			),

			// actual input
			components.CreateChatMessagesUser(
				components.ChatUserMessage{
					Content: components.CreateChatUserMessageContentStr(prompt),
					Role:    components.ChatUserMessageRoleUser,
				},
			),
		},
	}, nil)

	close(done)
	if err != nil {
		return fmt.Errorf("failed to send chat: %w", err)
	}
	fmt.Print("\r\033[K")

	// no response
	if res == nil || len(res.ChatResult.Choices) == 0 {
		return nil
	}

	// we need to unwrap it bc its nullable
	content := res.ChatResult.Choices[0].Message.Content
	if content.IsSet() {
		message, ok := content.Get()
		if !ok {
			return fmt.Errorf("failed to unwrap content")
		}

		if message.Str != nil {
			fmt.Println(*message.Str)
		}
	}

	return nil
}
