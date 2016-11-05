package service

import (
	"sync"

	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	"github.com/xh3b4sd/anna/service/spec"
)

// CollectionConfig represents the configuration used to create a new service
// collection object.
type CollectionConfig struct {
	// Dependencies.
	FSService          spec.FS
	IDService          spec.ID
	PermutationService spec.Permutation
	RandomService      spec.Random
}

// DefaultCollectionConfig provides a default configuration to create a new
// service collection object by best effort.
func DefaultCollectionConfig() CollectionConfig {
	newConfig := CollectionConfig{
		// Dependencies.
		FSService:          mem.MustNewService(),
		IDService:          id.MustNewService(),
		PermutationService: permutation.MustNewService(),
		RandomService:      random.MustNewService(),
	}

	return newConfig
}

// NewCollection creates a new configured service collection object.
func NewCollection(config CollectionConfig) (spec.ServiceCollection, error) {
	newCollection := &collection{
		CollectionConfig: config,

		ShutdownOnce: sync.Once{},
	}

	if newCollection.FSService == nil {
		return nil, maskAnyf(invalidConfigError, "file system service must not be empty")
	}
	if newCollection.IDService == nil {
		return nil, maskAnyf(invalidConfigError, "ID service must not be empty")
	}
	if newCollection.PermutationService == nil {
		return nil, maskAnyf(invalidConfigError, "permutation service must not be empty")
	}
	if newCollection.RandomService == nil {
		return nil, maskAnyf(invalidConfigError, "random service must not be empty")
	}

	return newCollection, nil
}

// MustNewCollection creates either a new default configured service collection
// object, or panics.
func MustNewCollection() spec.ServiceCollection {
	newCollection, err := NewCollection(DefaultCollectionConfig())
	if err != nil {
		panic(err)
	}

	return newCollection
}

type collection struct {
	CollectionConfig

	ShutdownOnce sync.Once
}

func (c *collection) FS() spec.FS {
	return c.FSService
}

func (c *collection) ID() spec.ID {
	return c.IDService
}

func (c *collection) Permutation() spec.Permutation {
	return c.PermutationService
}

func (c *collection) Random() spec.Random {
	return c.RandomService
}

func (c *collection) Shutdown() {
	c.ShutdownOnce.Do(func() {
		var wg sync.WaitGroup

		//wg.Add(1)
		//go func() {
		//	c.TODO().Shutdown()
		//	wg.Done()
		//}()

		wg.Wait()
	})
}
