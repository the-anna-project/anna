package spec

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// TextInput provides a communication channel to send information sequences
// back to the client.
type TextInput interface {
	// GetChannel returns a channel which is used to send text responses back to
	// the client.
	GetChannel() chan spec.TextInput
}
