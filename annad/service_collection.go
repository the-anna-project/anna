package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/spec"
)

func newServiceCollection() (spec.ServiceCollection, error) {
	fileSystemService, err := newFileSystemService()
	if err != nil {
		return nil, maskAny(err)
	}
	randomService, err := newRandomService()
	if err != nil {
		return nil, maskAny(err)
	}
	idService, err := newIDService(randomService)
	if err != nil {
		return nil, maskAny(err)
	}
	permutationService, err := newPermutationService()
	if err != nil {
		return nil, maskAny(err)
	}

	newCollectionConfig := service.DefaultCollectionConfig()
	newCollectionConfig.FSService = fileSystemService
	newCollectionConfig.IDService = idService
	newCollectionConfig.PermutationService = permutationService
	newCollectionConfig.RandomService = randomService
	newCollection, err := service.NewCollection(newCollectionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCollection, nil
}

// TODO make mem/os configurable
func newFileSystemService() (servicespec.FS, error) {
	newConfig := mem.DefaultConfig()
	newService, err := mem.New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newIDService(randomService servicespec.Random) (servicespec.ID, error) {
	newConfig := id.DefaultConfig()
	newConfig.RandomService = randomService
	newService, err := id.New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newPermutationService() (servicespec.Permutation, error) {
	newConfig := permutation.DefaultConfig()
	newService, err := permutation.New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newRandomService() (servicespec.Random, error) {
	newConfig := random.DefaultConfig()
	newConfig.BackoffFactory = func() spec.Backoff {
		return backoff.NewExponentialBackOff()
	}
	newService, err := random.New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}
