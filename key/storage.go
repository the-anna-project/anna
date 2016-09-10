package key

import (
	"fmt"
)

const (
	// CLGKeyFormat represents the format used to create storage keys for the CLG
	// scope. "s" stands for the scope, that is, the CLG scope. "o" stands for
	// the object requesting the key. "<key>" stands for the key-value pair
	// identifying the most specific part of the key being used.
	CLGKeyFormat = "s:clg:%s"
)

// TODO
//
//     add scope to CLG keys to identify current basics
//     e.g. output CLG reads from basic having most important certenty
//
//     add context to interface
//

// NewCLGKey returns a well configured key used to store and fetch data. Keys
// generated with NewCLGKey should only be used by objects related to the CLG
// scope. This can be e.g. each CLG. These objects generate and structure
// dynamic information used to learn.
func NewCLGKey(f string, v ...interface{}) string {
	return fmt.Sprintf(CLGKeyFormat, fmt.Sprintf(f, v...))
}

const (
	// NetKeyFormat represents the format used to create storage keys for the
	// neural network scope. "s" stands for the scope, that is, the network
	// scope. "o" stands for the object requesting the key. "<key>" stands for
	// the key-value pair identifying the most specific part of the key being
	// used.
	NetKeyFormat = "s:net:%s"
)

// NewNetKey returns a well configured key used to store and fetch data. Keys
// generated with NewNetKey should only be used by objects related to the
// neural network scope. This can be e.g. the neural network.
func NewNetKey(f string, v ...interface{}) string {
	return fmt.Sprintf(NetKeyFormat, fmt.Sprintf(f, v...))
}
