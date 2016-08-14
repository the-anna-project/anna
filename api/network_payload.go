package api

import (
	"reflect"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

// NetworkPayloadConfig represents the configuration used to create a new
// network payload object.
type NetworkPayloadConfig struct {
	// Settings.

	// Args represents the arguments intended to be used for the requested CLG
	// execution, or the output values being calculated during the requested CLG
	// execution. By convention this list must always contain a context as first
	// argument, otherwise the network payload is considered invalid. In this
	// case all calls to Validate would throw errors which can be asserted using
	// IsInvalidInterface. For more information about the context being used here
	// see https://godoc.org/golang.org/x/net/context.
	Args []reflect.Value

	// Destination represents the object ID of the CLG receiving the current
	// network payload.
	Destination spec.ObjectID

	// Sources represents the object IDs of the CLGs being involved providing the
	// current network payload. In fact, a network payload can only be sent by
	// one CLG. Reason for this being a list is the merge of network payloads
	// until a CLG's interface is satisfied. That way multiple CLG's can request
	// another CLG even if they are not able to satisfy the interface of the
	// requested CLG on their own. This gives the neural network an opportunity
	// to learn to combine CLGs as desired.
	Sources []spec.ObjectID
}

// DefaultNetworkPayloadConfig provides a default configuration to create a new
// network payload object by best effort.
func DefaultNetworkPayloadConfig() NetworkPayloadConfig {
	newConfig := NetworkPayloadConfig{
		Args:        nil,
		Destination: "",
		Sources:     nil,
	}

	return newConfig
}

// NewNetworkPayload creates a new configured network payload object.
func NewNetworkPayload(config NetworkPayloadConfig) (spec.NetworkPayload, error) {
	newNetworkPayload := &networkPayload{
		NetworkPayloadConfig: config,

		ID: id.MustNew(),
	}

	return newNetworkPayload, nil
}

type networkPayload struct {
	NetworkPayloadConfig

	ID spec.ObjectID
}

func (np *networkPayload) GetArgs() []reflect.Value {
	return np.Args
}

func (np *networkPayload) GetContext() (context.Context, error) {
	if len(np.Args) < 1 {
		return nil, maskAnyf(invalidConfigError, "arguments must have context")
	}
	ctx, ok := np.Args[0].Interface().(context.Context)
	if !ok {
		return nil, maskAnyf(invalidInterfaceError, "arguments must have context")
	}

	return ctx, nil
}

func (np *networkPayload) GetDestination() spec.ObjectID {
	return np.Destination
}

func (np *networkPayload) GetID() spec.ObjectID {
	return np.ID
}

func (np *networkPayload) GetSources() []spec.ObjectID {
	return np.Sources
}

func (np *networkPayload) Validate() error {
	// Check if the network payload has invalid properties.
	if np.Args == nil {
		return maskAnyf(invalidConfigError, "arguments must not be empty")
	}
	if np.Destination == "" {
		return maskAnyf(invalidConfigError, "destination must not be empty")
	}
	if len(np.Sources) < 1 {
		return maskAnyf(invalidConfigError, "sources must not be empty")
	}

	// Check if the network payload has an context as first argument.
	_, err := np.GetContext()
	if err != nil {
		return maskAny(err)
	}

	return nil
}
