package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"project/config"
)

type GPTRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type GPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func GenerateDescription(prompt string) (string, error) {
	gptRequest := GPTRequest{
		Model:       "text-davinci-003",
		Prompt:      prompt,
		MaxTokens:   150,
		Temperature: 0.7,
	}

	requestBody, err := json.Marshal(gptRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
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

	var gptResponse GPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&gptResponse); err != nil {
		return "", err
	}

	if len(gptResponse.Choices) > 0 {
		return gptResponse.Choices[0].Text, nil
	}

	return "", errors.New("no text returned by GPT API")
}
