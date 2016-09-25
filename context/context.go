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

const (
	behaviorIDKey    key = "behavior-id"
	clgTreeIDKey     key = "clg-tree-id"
	expectationIDKey key = "expectation"
	informationIDKey key = "information-id"
	sessionIDKey     key = "session-id"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// Settings.

	Context netcontext.Context

	// TODO we want to track the original input that was provided from the
	// outside. Further it would probably be interesting to also track the last 3
	// arguments of the current connection path.
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Context: netcontext.Background(),
	}

	return newConfig
}

// New creates a new configured context object.
func New(config Config) (spec.Context, error) {
	newContext := &context{
		Config: config,

		ID: string(id.MustNew()),
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
	// At first we create a new context with its own very unique ID, which will
	// not be cloned. All properties but the context ID must be cloned below.
	newContext := MustNew()

	// We prepare a new underlying context to have a fresh storage.
	newContext.(*context).Context = netcontext.Background()

	// Add the other information to the new underlying context.
	behaviorID, _ := c.GetBehaviorID()
	newContext.SetBehaviorID(behaviorID)
	clgTreeID, _ := c.GetCLGTreeID()
	newContext.SetCLGTreeID(clgTreeID)
	expectation, _ := c.GetExpectation()
	newContext.SetExpectation(expectation)
	informationID, _ := c.GetInformationID()
	newContext.SetInformationID(informationID)
	sessionID, _ := c.GetSessionID()
	newContext.SetSessionID(sessionID)

	return newContext
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

func (c *context) GetBehaviorID() (string, bool) {
	v, ok := c.Context.Value(behaviorIDKey).(string)
	return v, ok
}

func (c *context) GetCLGTreeID() (string, bool) {
	v, ok := c.Context.Value(clgTreeIDKey).(string)
	return v, ok
}

func (c *context) GetExpectation() (spec.Expectation, bool) {
	v, ok := c.Context.Value(expectationIDKey).(spec.Expectation)
	return v, ok
}

func (c *context) GetID() string {
	return c.ID
}

func (c *context) GetInformationID() (string, bool) {
	v, ok := c.Context.Value(informationIDKey).(string)
	return v, ok
}

func (c *context) GetSessionID() (string, bool) {
	v, ok := c.Context.Value(sessionIDKey).(string)
	return v, ok
}

func (c *context) SetBehaviorID(behaviorID string) {
	c.Context = netcontext.WithValue(c.Context, behaviorIDKey, behaviorID)
}

func (c *context) SetCLGTreeID(clgTreeID string) {
	c.Context = netcontext.WithValue(c.Context, clgTreeIDKey, clgTreeID)
}

func (c *context) SetExpectation(expectation spec.Expectation) {
	c.Context = netcontext.WithValue(c.Context, expectationIDKey, expectation)
}

func (c *context) SetInformationID(informationID string) {
	c.Context = netcontext.WithValue(c.Context, informationIDKey, informationID)
}

func (c *context) SetSessionID(sessionID string) {
	c.Context = netcontext.WithValue(c.Context, sessionIDKey, sessionID)
}

func (c *context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
