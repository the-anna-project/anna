// Package textoutput provides a simple service for sending text output
// responses.
package textoutput

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new text output service.
func New() servicespec.TextOutput {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	channel  chan objectspec.TextOutput
	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "text-output",
		"type": "service",
	}

	// Settings.
	s.channel = make(chan objectspec.TextOutput, 1000)
}

func (s *service) Channel() chan objectspec.TextOutput {
	return s.channel
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}
