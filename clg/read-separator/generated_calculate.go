package readseparator

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
	ObjectType spec.ObjectType = "read-separator-clg"
)

// Config represents the configuration used to create a new CLG object.
type Config struct {
	// Dependencies.
	FactoryCollection spec.FactoryCollection
	Log               spec.Log
	StorageCollection spec.StorageCollection
}

// DefaultConfig provides a default configuration to create a new CLG object by
// best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryCollection: factory.MustNewCollection(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),
	}

	return newConfig
}

// New creates a new configured CLG object.
func New(config Config) (spec.CLG, error) {
	newCLG := &clg{
		Config: config,
		ID:     id.MustNew(),
		Name:   "read-separator",
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
	Name string
	Type spec.ObjectType
}

func (c *clg) Factory() spec.FactoryCollection {
	return c.FactoryCollection
}

func (c *clg) GetCalculate() interface{} {
	return c.calculate
}

func (c *clg) GetName() string {
	return c.Name
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
