package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/runner"
	"google.golang.org/adk/v2/session"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/geminitool"
	"google.golang.org/genai"
)

func init() {
	rootCmd.AddCommand(aiCmd)
}

var aiCmd = &cobra.Command{
	Use:   "ai [input]",
	Short: "Interact with the model via CLI",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return callModel(cmd.Context(), args[0])
	},
}

func callModel(ctx context.Context, prompt string) error {
	model, err := gemini.NewModel(ctx, "gemini-3.1-flash-lite", &genai.ClientConfig{
		APIKey: os.Getenv("API_KEY"),
	})
	if err != nil {
		return fmt.Errorf("failed to create model: %w", err)
	}

	assistant, err := llmagent.New(llmagent.Config{
		Name:        "rb",
		Model:       model,
		Description: "A personal AI assistant for the terminal",
		Instruction: "You are a helpful technical assistant. Help with questions, debugging, logs, and code reviews.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create agent: %w", err)
	}

	sessionSvc := session.InMemoryService()

	_, err = sessionSvc.Create(ctx, &session.CreateRequest{
		AppName:   "rb",
		UserID:    "cli-user",
		SessionID: "session-1",
	})
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	r, err := runner.New(runner.Config{
		AppName:        "rb",
		Agent:          assistant,
		SessionService: sessionSvc,
	})
	if err != nil {
		return fmt.Errorf("failed to create runner: %w", err)
	}

	userContent := &genai.Content{
		Role: "user",
		Parts: []*genai.Part{
			genai.NewPartFromText(prompt),
		},
	}

	events := r.Run(ctx, "cli-user", "session-1", userContent, agent.RunConfig{})

	for ev, err := range events {
		if err != nil {
			return fmt.Errorf("error during run: %w", err)
		}

		if ev.Content == nil {
			continue
		}

		for _, part := range ev.Content.Parts {
			if part.Text != "" {
				fmt.Print(part.Text)
			}
		}
	}
	fmt.Println()

	return nil
}
