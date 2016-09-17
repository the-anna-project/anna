package spec

import (
	"golang.org/x/net/context"
)

// Context represents a container holding scope specific information.
type Context interface {
	// Clone returns an exact copy of the current Context. The only exception of
	// copied fields is the context ID, which must be unique for each context.
	Clone() Context

	context.Context

	// GetBehaviorID returns the behavior ID of the current Context. This behavior
	// ID represents the behavior currently being executed. That way CLGs can
	// identify themself.
	GetBehaviorID() string

	// GetCLGTreeID returns the CLG tree ID of the current Context.
	GetCLGTreeID() string

	// GetID returns the context's ID representing the very unique scope of its
	// own lifetime. This can be useful for e.g. gathering logs bound to one
	// request going through multiple independent sub-systems.
	GetID() string

	// GetInformationID returns the information ID of the current Context. This
	// information ID represents the information sequence of the original user
	// input.
	GetInformationID() string

	// GetSessionID returns the session ID of the current Context.
	GetSessionID() string

	// SetBehaviorID sets the given behavior ID to the current Context.
	SetBehaviorID(behaviorID string)

	// SetCLGTreeID sets the given CLG tree ID to the current Context.
	SetCLGTreeID(clgTreeID string)

	// SetInformationID sets the given information ID to the current Context.
	SetInformationID(informationID string)

	// SetSessionID sets the given session ID to the current Context.
	SetSessionID(sessionID string)
}
