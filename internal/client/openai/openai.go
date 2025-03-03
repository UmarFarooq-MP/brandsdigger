package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
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

// OpenAIMessageGen struct
type OpenAIMessageGen struct {
	ApiKey    string
	ApiUrl    string
	ModelName string
}

// GenerateNames - Generate 10 brand names using OpenAI API
func (oai OpenAIMessageGen) GenerateNames(message string) ([]string, error) {
	prompt := fmt.Sprintf(`I have a startup idea, and I need a brand name for it.
		Here is the idea:
		%s
		Your task:
		1. Generate 10 **unique, creative, and brandable** names that fit the startup's vision.
		2. Ensure each name is **short, easy to pronounce, and memorable**.
		3. **Check if the .com domain is available** for each name.
		4. Provide the names in a numbered list format.
		5. Please reply with only 10 names no extra words are required just simple 10 domain names.`, message)

	requestBody := ChatRequest{
		Model: oai.ModelName,
		Messages: []Message{
			{Role: "system", Content: "You are a creative AI that generates brand names with available domains."},
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", oai.ApiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+oai.ApiKey)

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
			cleanNames = append(cleanNames, cleanDomain(name))
		}
	}

	return cleanNames, nil
}

func cleanDomain(input string) string {
	// Define a regex to remove leading numbers and dots (e.g., "1. ")
	re := regexp.MustCompile(`^\d+\.\s*`)
	return re.ReplaceAllString(input, "")
}

// New - Creates a new OpenAIMessageGen instance
func New(apiKey, modelName string) *OpenAIMessageGen {
	return &OpenAIMessageGen{
		ApiUrl:    "https://api.openai.com/v1/chat/completions",
		ApiKey:    apiKey,
		ModelName: modelName,
	}
}
