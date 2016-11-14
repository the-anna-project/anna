// Package textinput provides a simple service for receiving text input
// requests.
package textinput

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new text input service.
func New() servicespec.TextInput {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	channel  chan objectspec.TextInput
	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "text-input",
		"type": "service",
	}

	s.channel = make(chan objectspec.TextInput, 1000)
}

func (s *service) Channel() chan objectspec.TextInput {
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
