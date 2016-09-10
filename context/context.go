// Package context implements spec.Context and provides a wrapper of
// golang.org/x/net/context with additional business logic related to the Anna
// project.
package context

import (
	"time"

	netcontext "golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

type key string

// TODO add canceler which can be passed through all contexted that are involved in responding to current request

const (
	behaviorIDKey key = "behavior-id"
	clgTreeIDKey  key = "clg-tree-id"
	sessionIDKey  key = "session-id"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// Settings.

	Context   netcontext.Context
	SessionID string

	// TODO we want to track the original input that was provided from the
	// outside. Further it would probably be interesting to also track the last 3
	// arguments of the current connection path.
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Context:   netcontext.Background(),
		SessionID: "",
	}

	return newConfig
}

// New creates a new configured context object.
func New(config Config) (spec.Context, error) {
	newContext := &context{
		Config: config,

		ID: string(id.MustNew()),
	}

	// If there is a session ID configured, we set it to the underlying context.
	// That way our standard configuration interface is obtained and the data
	// structures of the underlying implementation consistent.
	if config.SessionID != "" {
		newContext.SetSessionID(config.SessionID)
	}

	if newContext.Context == nil {
		return nil, maskAnyf(invalidConfigError, "context must not be empty")
	}

	return newContext, nil
}

// MustNew creates either a new default configured context object, or panics.
func MustNew() spec.Context {
	newContext, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newContext
}

type context struct {
	Config

	ID string
}

func (c *context) Clone() spec.Context {
	specContext := MustNew()

	specContext.(*context).Context = netcontext.Background()
	specContext.(*context).SessionID = c.GetSessionID()
	specContext.SetSessionID(c.GetSessionID())
	specContext.SetCLGTreeID(c.GetCLGTreeID())

	return specContext
}

func (c *context) Deadline() (time.Time, bool) {
	return c.Context.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *context) Err() error {
	return c.Context.Err()
}

func (c *context) GetBehaviorID() string {
	behaviorID, ok := c.Context.Value(behaviorIDKey).(string)
	if ok {
		return behaviorID
	}

	return ""
}

func (c *context) GetCLGTreeID() string {
	clgTreeID, ok := c.Context.Value(clgTreeIDKey).(string)
	if ok {
		return clgTreeID
	}

	return ""
}

func (c *context) GetID() string {
	return c.ID
}

func (c *context) GetSessionID() string {
	sessionID, ok := c.Context.Value(sessionIDKey).(string)
	if ok {
		return sessionID
	}

	return ""
}

func (c *context) SetBehaviorID(behaviorID string) {
	c.Context = netcontext.WithValue(c.Context, behaviorIDKey, behaviorID)
}

func (c *context) SetCLGTreeID(clgTreeID string) {
	c.Context = netcontext.WithValue(c.Context, clgTreeIDKey, clgTreeID)
}

func (c *context) SetSessionID(sessionID string) {
	c.Context = netcontext.WithValue(c.Context, sessionIDKey, sessionID)
}

func (c *context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
