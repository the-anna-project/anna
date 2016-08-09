package spec

import (
	"reflect"
)

// NetworkPayload represents the request provided to a CLG to ask it to do some
// work.
type NetworkPayload struct {
	// Args represents the arguments intended to be used for the requested CLG
	// execution, or the output values being calculated during the requested CLG
	// execution.
	Args []reflect.Value

	// Destination represents the ID of the CLG that receives the message.
	Destination ObjectID

	// Sources represents the IDs of the CLGs that sent the message.
	Sources []ObjectID
}

// Network provides a neural network based on dynamic and self improving CLG
// execution. The network provides input and output channels. When input is
// received it is injected into the neural communication. The following neural
// activity calculates output which is streamed through the output channel back
// to the requestor.
//
// Network
//
//     At the very beginning there is the neural network. It initializes all known
//     CLGs. Every CLG has an input and an output channel.
//
// Network.Listen
//
//     The network listens on each CLG input channel using Listen.
//
// Network.Send
//
//     The Send method is used to emit the Input CLG by providing the impulses
//     input. Once send, the input is submitted to the neural network.
//
// Network.Execute
//
//     As stated above, Listen was initialized to wait for inputs of each CLG.
//     Now the Input CLG received some input. Thus it is executed using Execute
//     and the provided input.
//
// Network.Activate
//
//     Each CLG that is executed needs to decide if it wants to be activated.
//     This happens using the Activate method. To make this decision the given
//     input, the CLGs connections and behaviour properties are considered.
//
// Network.Calculate
//
//     Once activated, a CLG executes its actual implemented behaviour using
//     Calculate. This behaviour can be anything. It is up to the CLG.
//
// Network.Forward
//
//     After the CLGs calculation it can decide what to do next. Like Activate,
//     it is up to the CLG if it forwards signals to further CLGs. E.g. a CLG
//     might or might not forward its calculated results to one or more CLGs.
//     All this depends on its inputs, calculated outputs, CLG connections and
//     behaviour properties.
//
// Network.Receive
//
//     In the Trigger method, the Input CLG takes the Impulse's input and
//     causes the neural network to trigger. There the Output CLG is asked to
//     provide some output. Trigger waits until the Output CLG returns
//     something using Receive. In case Receive returns in behalf of the Output
//     CLG at some point in time, the Impulse found a way from the Input CLG,
//     through the neural network, up to the Output CLG. Some expectation
//     matching might be required, if provided with the request. If the
//     expectation, if any, matches, the calculated output is returned. If it
//     does not match, the procedure starts again with Send. Then the
//     calculated output is used as input for the next iteration.
//
type Network interface {
	Activate(clgID ObjectID, inputs []reflect.Value) (bool, error)

	// Boot initializes and starts the whole network like booting a machine. The
	// call to Boot blocks until the network is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()

	Calculate(clgID ObjectID, inputs []reflect.Value) ([]reflect.Value, error)

	Execute(clgID ObjectID, requests []NetworkPayload) error

	Forward(clgID ObjectID, inputs, outputs []reflect.Value) error

	Listen()

	Object

	Receive(clgID ObjectID) (NetworkPayload, error)

	Send(request NetworkPayload) error

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
