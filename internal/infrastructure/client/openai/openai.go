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
func (oai OpenAIMessageGen) GenerateNames(idea string) ([]string, error) {

	// STEP 1 — Simple sanity validation before calling OpenAI
	if strings.TrimSpace(idea) == "" {
		return nil, fmt.Errorf("invalid input: empty brand idea")
	}

	// STEP 2 — A clean, JSON-safe request prompt
	prompt := fmt.Sprintf(`
I have an idea for a startup or brand:

"%s"

### Your Rules:
1. First evaluate the idea.
2. If it is NOT a startup/business idea (empty, greeting, or unclear),
   respond EXACTLY with:
   { "error": "invalid_input" }

3. If it IS valid:
   - Generate 10 SHORT, BRANDABLE .com domain names.
   - NO sentences. NO explanation. NO description.
   - Only domain names (like "FlowZen.com").
   - Output ONLY JSON in this format:

{
  "names": ["name1.com", "name2.com", ...]
}
`, idea)

	requestBody := ChatRequest{
		Model: oai.ModelName,
		Messages: []Message{
			{Role: "system", Content: "You generate ONLY JSON and follow instructions strictly."},
			{Role: "user", Content: prompt},
		},
	}

	// STEP 3 — Marshal request
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}

	// STEP 4 — Build HTTP request
	req, err := http.NewRequest("POST", oai.ApiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+oai.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI: %w", err)
	}
	defer resp.Body.Close()

	// STEP 5 — Must be HTTP 200
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("openai error: %s → %s", resp.Status, string(body))
	}

	// STEP 6 — Decode OpenAI response
	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return nil, fmt.Errorf("failed decoding response JSON: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return nil, fmt.Errorf("no response from openai")
	}

	content := chatResp.Choices[0].Message.Content

	// STEP 7 — Detect invalid_input JSON
	if strings.Contains(content, `"error":`) && strings.Contains(content, `invalid_input`) {
		return nil, fmt.Errorf("invalid brand idea")
	}

	// STEP 8 — Extract JSON safely
	type nameResponse struct {
		Names []string `json:"names"`
	}

	var nr nameResponse
	if err := json.Unmarshal([]byte(content), &nr); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAI JSON: %w\nRaw: %s", err, content)
	}

	if len(nr.Names) == 0 {
		return nil, fmt.Errorf("no names returned")
	}

	// STEP 9 — Final cleanup (remove weird whitespace)
	var cleaned []string
	for _, name := range nr.Names {
		cleaned = append(cleaned, strings.TrimSpace(name))
	}

	return cleaned, nil
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
