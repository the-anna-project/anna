package main

import (
	"os"

	kitlog "github.com/go-kit/kit/log"

	"github.com/the-anna-project/collection"
	"github.com/the-anna-project/fs/memory"
	"github.com/the-anna-project/id"
	inputcollection "github.com/the-anna-project/input/collection"
	textinputservice "github.com/the-anna-project/input/service/text"
	"github.com/the-anna-project/log"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/textoutput"
)

func (a *annactl) newServiceCollection() servicespec.ServiceCollection {
	// Set.
	collection := collection.New()

	collection.SetFSService(a.newFSService())
	collection.SetIDService(a.newIDService())
	collection.SetInputCollection(a.newInputCollection())
	collection.SetLogService(a.newLogService())
	collection.SetPermutationService(a.newPermutationService())
	collection.SetTextOutputService(a.newTextOutputService())

	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Input().Text().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.TextOutput().SetServiceCollection(collection)

	return collection
}

// TODO make mem/os configurable
func (a *annactl) newFSService() servicespec.FSService {
	return memory.New()
}

func (a *annactl) newIDService() servicespec.IDService {
	return id.New()
}

func (a *annactl) newInputCollection() servicespec.InputCollection {
	newCollection := inputcollection.New()

	newCollection.SetTextService(textinputservice.New())

	return newCollection
}

func (a *annactl) newLogService() servicespec.LogService {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (a *annactl) newPermutationService() servicespec.PermutationService {
	return permutation.New()
}

func (a *annactl) newTextOutputService() servicespec.TextOutputService {
	return textoutput.New()
}
