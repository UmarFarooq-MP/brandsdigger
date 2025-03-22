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
	prompt := fmt.Sprintf(`I have an idea for a **new startup or brand**, and I need a name for it.
	**Here is the idea:** %s  

	### Your Task:
		1. Validate the input. If the provided idea is **empty, a greeting, or unrelated to a startup/brand**, respond with:  
		   **"Invalid input. Please provide a clear startup or brand idea."** and do not proceed further.  
		2. Generate **10 unique, creative, and brandable** names that align with the given idea.
		3. Ensure each name is **short, easy to pronounce, and memorable**.
		4. **Check if the .com domain is available** for each name.
		5. Present the names in a **simple numbered list** (1-10) with no extra words or explanations.

	**Output format:**  
		- If the input is valid: return **only** the 10 domain names.  
		- If the input is invalid: return **"Invalid input. Please provide a clear startup or brand idea."**`, message)

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
