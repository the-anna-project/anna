package api

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/spec"
)

// TextRequestConfig represents the configuration used to create a new text
// response object.
type TextRequestConfig struct {
	// Settings.

	// ExpectationRequest represents the expectation object. This is used to
	// match against the calculated output. In case there is an expectation
	// given, the neural network tries to calculate an output until it generated
	// one that matches the given expectation.
	ExpectationRequest ExpectationRequest `json:"expectation,omitempty"`

	// Input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	Input string `json:"input"`

	// SessionID represents the session the current text request is associated
	// with. This is provided to differentiate streams between different users.
	SessionID string `json:"session_id,omitempty"`
}

// DefaultTextRequestConfig provides a default configuration to create a new
// text request object by best effort.
func DefaultTextRequestConfig() TextRequestConfig {
	newConfig := TextRequestConfig{
		ExpectationRequest: ExpectationRequest{},
		Input:              "",
		SessionID:          "",
	}

	return newConfig
}

// NewTextRequest creates a new configured text request object.
func NewTextRequest(config TextRequestConfig) (spec.TextRequest, error) {
	newTextRequest := &textRequest{
		TextRequestConfig: config,
	}

	return newTextRequest, nil
}

// NewEmptyTextRequest simply returns an empty, maybe invalid, text request
// object. This should only be used for things like unmarshaling.
func NewEmptyTextRequest() spec.TextRequest {
	return &textRequest{}
}

type textRequest struct {
	TextRequestConfig
}

func (tr *textRequest) GetInput() string {
	return tr.Input
}

func (tr *textRequest) IsEmpty() bool {
	return tr.Input == "" || tr.SessionID == ""
}

// textRequestClone is for making use of the stdlib json implementation. The
// textRequest object implements its own marshaler and unmarshaler but only to
// provide json implementations for spec.TextRequest. Note, not redirecting
// the type will cause infinite recursion.
type textRequestClone textRequest

func (tr *textRequest) MarshalJSON() ([]byte, error) {
	newTextRequest := textRequestClone(*tr)

	raw, err := json.Marshal(newTextRequest)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (tr *textRequest) UnmarshalJSON(b []byte) error {
	newTextRequest := textRequestClone{}

	err := json.Unmarshal(b, &newTextRequest)
	if err != nil {
		return maskAny(err)
	}

	tr.ExpectationRequest = newTextRequest.ExpectationRequest
	tr.Input = newTextRequest.Input
	tr.SessionID = newTextRequest.SessionID

	return nil
}
