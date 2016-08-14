package text

import (
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeTextInterface represents the object type of the text interface
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeTextInterface spec.ObjectType = "text-interface"
)

// InterfaceConfig represents the configuration used to create a new text
// interface object.
type InterfaceConfig struct {
	Log        spec.Log
	TextInput  chan spec.TextRequest
	TextOutput chan spec.TextResponse
}

// DefaultInterfaceConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultInterfaceConfig() InterfaceConfig {
	newConfig := InterfaceConfig{
		Log:        log.NewLog(log.DefaultConfig()),
		TextInput:  make(chan spec.TextRequest, 1000),
		TextOutput: make(chan spec.TextResponse, 1000),
	}

	return newConfig
}

// NewInterface creates a new configured text interface object.
func NewInterface(config InterfaceConfig) (spec.TextInterface, error) {
	newInterface := &tinterface{
		InterfaceConfig: config,
		ID:              id.MustNew(),
		Mutex:           sync.Mutex{},
		Type:            spec.ObjectType(ObjectTypeTextInterface),
	}

	if newInterface.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	newInterface.Log.Register(newInterface.GetType())

	return newInterface, nil
}

// tinterface is not named interface because this is a reserved key in golang.
type tinterface struct {
	InterfaceConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *tinterface) StreamText(ctx context.Context, in chan spec.TextRequest, out chan spec.TextResponse) error {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call StreamText")

	fail := make(chan error, 1000)

	// Start processing the text request through the text input channel.
	go func() {
		for {
			select {
			case <-ctx.Done():
				fail <- maskAny(ctx.Err())
				return
			case textRequest := <-in:
				i.TextInput <- textRequest
			}
		}
	}()

	for {
		select {
		case err := <-fail:
			return maskAny(err)
		case <-ctx.Done():
			return maskAny(ctx.Err())
		case textResponse := <-i.TextOutput:
			out <- textResponse
		}
	}
}
