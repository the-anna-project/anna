package spec

import (
	"reflect"

	"golang.org/x/net/context"
)

// NetworkPayload represents the data container carried around within the
// neural network.
type NetworkPayload interface {
	// GetArgs returns the arguments of the current network payload.
	GetArgs() []reflect.Value

	// GetContext returns the context the current network payload holds as first
	// argument. If no context can be found, an error is returned.
	GetContext() (context.Context, error)

	// GetArgs returns the destination of the current network payload.
	GetDestination() ObjectID

	// GetID returns the object ID of the current network payload.
	GetID() ObjectID

	// GetArgs returns the sources of the current network payload.
	GetSources() []ObjectID

	// Validate throws an error if the current network payload is not valid. An
	// network payload is not valid if it is empty, or if it does not satisfy the
	// convention of the CLG interface to have a proper context as first input
	// and output parameter. For more information about the context being passed
	// through see https://godoc.org/golang.org/x/net/context.
	Validate() error
}

// Network provides a neural network based on dynamic and self improving CLG
// execution. The network provides input and output channels. When input is
// received it is injected into the neural communication. The following neural
// activity calculates outputs which are streamed through the output channel
// back to the requestor.
type Network interface {
	// Activate decides if the requested CLG should be activated. To make this
	// decision the given network payload is considered.
	// TODO explain interface better
	Activate(clgID spec.ObjectID, payload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error)

	// Boot initializes and starts the whole network like booting a machine. The
	// call to Boot blocks until the network is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()

	// Calculate executes the activated CLG and invokes its actual implemented
	// behaviour. This behaviour can be anything. It is up to the CLG what it
	// does with the provided NetworkPayload.
	Calculate(clgID ObjectID, payload NetworkPayload) (NetworkPayload, error)

	// Forward is triggered after the CLGs calculation. Here is decided what to
	// do next. Like Activate, it is up to the CLG if it forwards signals to
	// further CLGs. E.g. a CLG may or may not forward its calculated results to
	// one or more CLGs. All this depends on the information provided by the
	// given network payload, the CLG's connections and its therefore resulting
	// behaviour properties.
	Forward(clgID ObjectID, payload NetworkPayload) error

	// Listen makes the network listen on requests from the outside. Here each
	// CLG input channel is managed. This models Listen as kind of cortex in
	// which impulses are dispatched into all possible direction and finally flow
	// back again. Listen only returns an error in case the initialization of all
	// listeners failed. Errors during processing of the neural network will be
	// logged to the provided logger.
	Listen() error

	Object

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
