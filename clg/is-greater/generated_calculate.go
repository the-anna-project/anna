package isgreater

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the clg package. There is the go generate statement placed to invoke clggen.

import (
	"reflect"

	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	// ObjectType represents the object type of the CLG object. This is used e.g.
	// to register itself to the logger.
	ObjectType spec.ObjectType = "is-greater-clg"
)

// Config represents the configuration used to create a new CLG object.
type Config struct {
	// Dependencies.
	FactoryCollection spec.FactoryCollection
	Log               spec.Log
	StorageCollection spec.StorageCollection

	// Settings.
	InputChannel chan spec.NetworkPayload
}

// DefaultConfig provides a default configuration to create a new CLG object by
// best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryCollection: factory.MustNewCollection(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),

		// Settings.
		InputChannel: make(chan spec.NetworkPayload, 1000),
	}

	return newConfig
}

// New creates a new configured CLG object.
func New(config Config) (spec.CLG, error) {
	newCLG := &clg{
		Config: config,
		ID:     id.MustNew(),
		Type:   ObjectType,
	}

	// Dependencies.
	if newCLG.FactoryCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newCLG.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newCLG.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	// Settings.
	if newCLG.InputChannel == nil {
		return nil, maskAnyf(invalidConfigError, "input channel must not be empty")
	}

	newCLG.Log.Register(newCLG.GetType())

	return newCLG, nil
}

// MustNew creates either a new default configured CLG object, or panics.
func MustNew() spec.CLG {
	newCLG, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newCLG
}

type clg struct {
	Config

	ID   spec.ObjectID
	Type spec.ObjectType
}

func (c *clg) Calculate(payload spec.NetworkPayload) (spec.NetworkPayload, error) {
	outputs := reflect.ValueOf(c.calculate).Call(payload.GetArgs())

	payload.SetArgs(outputs)

	return payload, nil
}

func (c *clg) Factory() spec.FactoryCollection {
	return c.FactoryCollection
}

func (c *clg) GetName() string {
	return "is-greater"
}

func (c *clg) GetInputChannel() chan spec.NetworkPayload {
	return c.InputChannel
}

func (c *clg) GetInputTypes() []reflect.Type {
	t := reflect.TypeOf(c.calculate)

	var inputType []reflect.Type

	for i := 0; i < t.NumIn(); i++ {
		inputType = append(inputType, t.In(i))
	}

	return inputType
}

func (c *clg) SetFactoryCollection(factoryCollection spec.FactoryCollection) {
	c.FactoryCollection = factoryCollection
}

func (c *clg) SetLog(log spec.Log) {
	c.Log = log
}

func (c *clg) SetStorageCollection(storageCollection spec.StorageCollection) {
	c.StorageCollection = storageCollection
}

func (c *clg) Storage() spec.StorageCollection {
	return c.StorageCollection
}
