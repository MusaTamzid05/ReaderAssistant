package llm

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io"
	"strings"

)

type OllamaRequest struct {
	Model string   `json:"model"`
	Prompt string  `json:"prompt"`
	Stream bool    `json:"stream"`
	Think bool    `json:"think"`
	Options map[string]any `json:"options,omitempty"`

}


type OllamaResponse struct {
	Response string `json:"response"`
	Thinking string `json:"thinking"`
}

type Ollama struct {
	generationToken int // token that are use to generate result
	                   // if thinking is true then this token is
					   // also used for thinking

}

func MakeOllama() Ollama {
	return Ollama{generationToken: 8192}
}


func (o Ollama) Ask(question, document string) (string, error) {
	prompt := fmt.Sprintf(`
	Answer the question using only the provided document.
	If the answer is not available, say you don't know.

	DOCUMENT:
	%s

	QUESTION:
	%s
	`, document, question)

	requestBody := OllamaRequest {
		Model: "qwen3.5:9b",
		Prompt: prompt,
		Stream: false,
		Think: false,
		Options: map[string]any {
			"num_predict": o.generationToken,
		},
	}
	data, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama error %s", string(body))
	}

	var result OllamaResponse

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)

	if err != nil {
		return "", err
	}


	return strings.TrimSpace(result.Response), nil

}


