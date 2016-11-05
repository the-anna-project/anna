package api

import (
	"github.com/xh3b4sd/anna/service/spec"
)

// TextResponseConfig represents the configuration used to create a new text
// response object.
type TextResponseConfig struct {
	// Settings.

	// Output represents the output being calculated by the neural network.
	Output string `json:"output"`
}

// DefaultTextResponseConfig provides a default configuration to create a new
// text response object by best effort.
func DefaultTextResponseConfig() TextResponseConfig {
	newConfig := TextResponseConfig{
		Output: "",
	}

	return newConfig
}

// NewTextResponse creates a new configured text response object.
func NewTextResponse(config TextResponseConfig) (spec.TextResponse, error) {
	newTextResponse := &textResponse{
		TextResponseConfig: config,
	}

	return newTextResponse, nil
}

type textResponse struct {
	TextResponseConfig
}

func (tr *textResponse) GetOutput() string {
	return tr.Output
}
