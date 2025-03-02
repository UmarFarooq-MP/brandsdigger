package deepseek

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Request structure
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message structure
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response structure
type ChatResponse struct {
	Choices []Choice `json:"choices"`
}

// Choice structure
type Choice struct {
	Message Message `json:"message"`
}

type DeepSeekMessageGen struct {
	ApiKey    string
	ApiUrl    string
	ModelName string
}

// GenerateNames Generate 10 brand names with domain availability check prompt
func (ds DeepSeekMessageGen) GenerateNames(message string) ([]string, error) {
	prompt := fmt.Sprintf(`I have a startup idea, and I need a brand name for it.
		Here is the idea:
		%s
		Your task:
		1. Generate 10 **unique, creative, and brandable** names that fit the startup's vision.
		2. Ensure each name is **short, easy to pronounce, and memorable**.
		3. **Check if the .com domain is available** for each name.
		4. Provide the names in a numbered list format.`, message)

	requestBody := ChatRequest{
		Model: ds.ModelName,
		Messages: []Message{
			{Role: "system", Content: "You are a creative AI that generates brand names with available domains."},
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", ds.ApiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ds.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned error: %s, Response: %s", resp.Status, string(body))
	}

	var chatResp ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&chatResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response JSON: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return nil, fmt.Errorf("no brand names returned")
	}

	// Extract brand names from response
	names := strings.Split(chatResp.Choices[0].Message.Content, "\n")
	var cleanNames []string
	for _, name := range names {
		name = strings.TrimSpace(name)
		if name != "" {
			cleanNames = append(cleanNames, name)
		}
	}

	return cleanNames, nil
}

func New(apiKey, apiUrl, modelName string) DeepSeekMessageGen {
	return DeepSeekMessageGen{
		ApiUrl:    apiUrl,
		ApiKey:    apiKey,
		ModelName: modelName,
	}
}
