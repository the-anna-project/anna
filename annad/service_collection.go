package main

import (
	"os"

	"github.com/cenk/backoff"
	kitlog "github.com/go-kit/kit/log"

	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/activator"
	"github.com/xh3b4sd/anna/service/feature"
	"github.com/xh3b4sd/anna/service/forwarder"
	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/instrumentor/prometheus"
	"github.com/xh3b4sd/anna/service/log"
	"github.com/xh3b4sd/anna/service/network"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	"github.com/xh3b4sd/anna/service/server"
	"github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/service/storage"
	"github.com/xh3b4sd/anna/service/storage/memory"
	"github.com/xh3b4sd/anna/service/storage/redis"
	"github.com/xh3b4sd/anna/service/textendpoint"
	"github.com/xh3b4sd/anna/service/textinput"
	"github.com/xh3b4sd/anna/service/textoutput"
	"github.com/xh3b4sd/anna/service/tracker"
)

func (a *annad) newServiceCollection() spec.Collection {
	// Set.
	collection := service.NewCollection()

	collection.SetActivator(a.newActivatorService())
	collection.SetFeature(a.newFeatureService())
	collection.SetForwarder(a.newForwarderService())
	collection.SetFS(a.newFSService())
	collection.SetID(a.newIDService())
	collection.SetInstrumentor(a.newInstrumentorService())
	collection.SetLog(a.newLogService())
	collection.SetNetwork(a.newNetworkService())
	collection.SetPermutation(a.newPermutationService())
	collection.SetRandom(a.newRandomService())
	collection.SetServer(a.newServerService())
	collection.SetStorageCollection(a.newStorageCollection())
	collection.SetTextEndpoint(a.newTextEndpointService())
	collection.SetTextInput(a.newTextInputService())
	collection.SetTextOutput(a.newTextOutputService())
	collection.SetTracker(a.newTrackerService())

	collection.Activator().SetServiceCollection(collection)
	collection.Feature().SetServiceCollection(collection)
	collection.Forwarder().SetServiceCollection(collection)
	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Instrumentor().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Network().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.Random().SetServiceCollection(collection)
	collection.Server().SetServiceCollection(collection)
	collection.Storage().Connection().SetServiceCollection(collection)
	collection.Storage().Feature().SetServiceCollection(collection)
	collection.Storage().General().SetServiceCollection(collection)
	collection.TextEndpoint().SetServiceCollection(collection)
	collection.TextInput().SetServiceCollection(collection)
	collection.TextOutput().SetServiceCollection(collection)
	collection.Tracker().SetServiceCollection(collection)

	// Validate.
	panicOnError(collection.Validate())

	panicOnError(collection.Activator().Validate())
	panicOnError(collection.Feature().Validate())
	panicOnError(collection.Forwarder().Validate())
	panicOnError(collection.FS().Validate())
	panicOnError(collection.ID().Validate())
	panicOnError(collection.Instrumentor().Validate())
	panicOnError(collection.Log().Validate())
	panicOnError(collection.Network().Validate())
	panicOnError(collection.Permutation().Validate())
	panicOnError(collection.Random().Validate())
	panicOnError(collection.Server().Validate())
	panicOnError(collection.Storage().Validate())
	panicOnError(collection.Storage().Connection().Validate())
	panicOnError(collection.Storage().Feature().Validate())
	panicOnError(collection.Storage().General().Validate())
	panicOnError(collection.TextEndpoint().Validate())
	panicOnError(collection.TextInput().Validate())
	panicOnError(collection.TextOutput().Validate())
	panicOnError(collection.Tracker().Validate())

	// Configure.
	panicOnError(collection.Configure())

	panicOnError(collection.Activator().Configure())
	panicOnError(collection.Feature().Configure())
	panicOnError(collection.Forwarder().Configure())
	panicOnError(collection.FS().Configure())
	panicOnError(collection.ID().Configure())
	panicOnError(collection.Instrumentor().Configure())
	panicOnError(collection.Log().Configure())
	panicOnError(collection.Network().Configure())
	panicOnError(collection.Permutation().Configure())
	panicOnError(collection.Random().Configure())
	panicOnError(collection.Server().Configure())
	panicOnError(collection.Storage().Configure())
	panicOnError(collection.Storage().Connection().Configure())
	panicOnError(collection.Storage().Feature().Configure())
	panicOnError(collection.Storage().General().Configure())
	panicOnError(collection.TextEndpoint().Configure())
	panicOnError(collection.TextInput().Configure())
	panicOnError(collection.TextOutput().Configure())
	panicOnError(collection.Tracker().Configure())

	return collection
}

func (a *annad) newActivatorService() spec.Activator {
	return activator.New()
}

func (a *annad) newBackoffFactory() func() spec.Backoff {
	return func() spec.Backoff {
		return backoff.NewExponentialBackOff()
	}
}

func (a *annad) newFeatureService() spec.Feature {
	return feature.New()
}

func (a *annad) newForwarderService() spec.Forwarder {
	return forwarder.New()
}

// TODO make mem/os configurable
func (a *annad) newFSService() spec.FS {
	return mem.New()
}

func (a *annad) newIDService() spec.ID {
	return id.New()
}

func (a *annad) newInstrumentorService() spec.Instrumentor {
	return prometheus.New()
}

func (a *annad) newLogService() spec.Log {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (a *annad) newNetworkService() spec.Network {
	return network.New()
}

func (a *annad) newPermutationService() spec.Permutation {
	return permutation.New()
}

func (a *annad) newRandomService() spec.Random {
	newService := random.New()

	newService.SetBackoffFactory(a.newBackoffFactory())

	return newService
}

func (a *annad) newServerService() spec.Server {
	newService := server.New()

	newService.SetHTTPAddress(a.flags.HTTPAddr)

	return newService
}

func (a *annad) newStorageCollection() spec.StorageCollection {
	newCollection := storage.NewCollection()

	switch a.flags.StorageType {
	case "redis":
		connectionService := redis.New()
		connectionService.SetAddress(a.flags.RedisConnectionStorageAddr)
		connectionService.SetBackoffFactory(a.newBackoffFactory())
		connectionService.SetPrefix(a.flags.RedisConnectionStoragePrefix)
		newCollection.SetConnection(connectionService)

		featureService := redis.New()
		featureService.SetAddress(a.flags.RedisFeatureStorageAddr)
		featureService.SetBackoffFactory(a.newBackoffFactory())
		featureService.SetPrefix(a.flags.RedisFeatureStoragePrefix)
		newCollection.SetConnection(featureService)

		generalService := redis.New()
		generalService.SetAddress(a.flags.RedisGeneralStorageAddr)
		generalService.SetBackoffFactory(a.newBackoffFactory())
		generalService.SetPrefix(a.flags.RedisGeneralStoragePrefix)
		newCollection.SetConnection(generalService)
	case "memory":
		newCollection.SetConnection(memory.New())
		newCollection.SetFeature(memory.New())
		newCollection.SetGeneral(memory.New())
	default:
		panic(maskAnyf(invalidStorageTypeError, "%s", a.flags.StorageType))
	}

	return newCollection
}

func (a *annad) newTextEndpointService() spec.TextEndpoint {
	newService := textendpoint.New()

	newService.SetGRPCAddress(a.flags.GRPCAddr)

	return newService
}

func (a *annad) newTextInputService() spec.TextInput {
	return textinput.New()
}

func (a *annad) newTextOutputService() spec.TextOutput {
	return textoutput.New()
}

func (a *annad) newTrackerService() spec.Tracker {
	return tracker.New()
}
