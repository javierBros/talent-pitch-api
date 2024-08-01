package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"project/config"
)

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTChatRequest struct {
	Model       string       `json:"model"`
	Messages    []GPTMessage `json:"messages"`
	Temperature float64      `json:"temperature"`
	MaxTokens   int          `json:"max_tokens"`
	TopP        float64      `json:"top_p"`
}

type GPTChatResponse struct {
	Choices []struct {
		Message GPTMessage `json:"message"`
	} `json:"choices"`
}

func GenerateDescription(prompt string) (string, error) {
	gptRequest := GPTChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []GPTMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.8,
		MaxTokens:   64,
		TopP:        1,
	}

	requestBody, err := json.Marshal(gptRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.AppConfig.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get response from GPT API")
	}

	var gptResponse GPTChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&gptResponse); err != nil {
		return "", err
	}

	if len(gptResponse.Choices) > 0 {
		return gptResponse.Choices[0].Message.Content, nil
	}

	return "", errors.New("no text returned by GPT API")
}
