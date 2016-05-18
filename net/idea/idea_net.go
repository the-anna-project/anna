// Package ideanet implements spec.Network to provide functionality to bring in
// creative ideas into the inout processing and output creation.
package ideanet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeIdeaNet represents the object type of the idea network object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeIdeaNet spec.ObjectType = "idea-net"
)

// Config represents the configuration used to create a new idea network
// object.
type Config struct {
	Log     spec.Log
	Storage spec.Storage

	EvalNet  spec.Network
	ExecNet  spec.Network
	PatNet   spec.Network
	PredNet  spec.Network
	StratNet spec.Network
}

// DefaultConfig provides a default configuration to create a new idea network
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewIdeaNet creates a new configured idea network object.
func NewIdeaNet(config Config) (spec.Network, error) {
	newNet := &ideaNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeIdeaNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type ideaNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (in *ideaNet) Boot() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Boot")

	in.BootOnce.Do(func() {
	})
}

func (in *ideaNet) Shutdown() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Shutdown")

	in.ShutdownOnce.Do(func() {
	})
}

func (in *ideaNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	var err error
	for {
		imp, err = in.StratNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = in.PredNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = in.ExecNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = in.EvalNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}

		break
	}

	// Note that the impulse returned here is not actually the same as received
	// at the beginning of the call, but was manipulated during its walk through
	// the networks.
	return imp, nil
}
