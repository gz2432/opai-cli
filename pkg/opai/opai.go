package opai

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	viper "github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

const (
	DefaultMaxTokens   = 1024
	DefaultTemperature = 0.1
)

func Complete(prompt string) (string, error) {
	defer timeTrack(time.Now(), "Complete")
	client := gogpt.NewClient(viper.GetString("token"))
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:       viper.GetString("model"),
		MaxTokens:   viper.GetInt("max-tokens"),
		Temperature: float32(viper.GetFloat64("temperature")),
		Prompt:      prompt,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return strings.Trim(resp.Choices[0].Text, "\n"), nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
