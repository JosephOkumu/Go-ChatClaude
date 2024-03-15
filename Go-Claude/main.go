package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/liushuangls/go-anthropic"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	client := anthropic.NewClient(apiKey)
	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaudeInstant1Dot2,
		Messages: []anthropic.Message{
			anthropic.NewUserTextMessage("Give me three of your names?"),
		},
		MaxTokens: 500,
	})
	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			fmt.Printf("Messages error, type: %s, message: %s", e.Type, e.Message)
		} else {
			fmt.Printf("Messages error: %v\n", err)
		}
		return
	}
	fmt.Println(resp.Content[0].Text)
}
