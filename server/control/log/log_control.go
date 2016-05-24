// Package logcontrol implements spec.LogControl to interactively configure
// what is being logged through the network API.
package logcontrol

import (
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeLogControl represents the object type of the log control object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeLogControl spec.ObjectType = "log-control"
)

// Config represents the configuration used to create a new log control object.
type Config struct {
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new log control
// object by best effort.
func DefaultConfig() Config {
	return Config{
		Log: log.NewLog(log.DefaultConfig()),
	}
}

// NewLogControl creates a new configured log control object.
func NewLogControl(config Config) spec.LogControl {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newControl := &logControl{
		Config: config,
		ID:     newID,
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(ObjectTypeLogControl),
	}

	newControl.Log.Register(newControl.GetType())

	return newControl
}

type logControl struct {
	Config

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (lc *logControl) ResetLevels(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetLevels")

	err := lc.Log.ResetLevels()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) ResetObjects(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetObjects")

	err := lc.Log.ResetObjects()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) ResetVerbosity(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetVerbosity")

	err := lc.Log.ResetVerbosity()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) SetLevels(ctx context.Context, levels string) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetLevels")

	err := lc.Log.SetLevels(levels)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) SetObjects(ctx context.Context, objectTypes string) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetObjects")

	err := lc.Log.SetObjects(objectTypes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) SetVerbosity(ctx context.Context, verbosity int) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetVerbosity")

	err := lc.Log.SetVerbosity(verbosity)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
