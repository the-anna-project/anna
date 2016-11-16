package service

import (
	"sync"

	servicespec "github.com/the-anna-project/spec/service"
)

// NewCollection creates a new service collection.
func NewCollection() servicespec.ServiceCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	activatorService    servicespec.ActivatorService
	connectionService   servicespec.ConnectionService
	endpointCollection  servicespec.EndpointCollection
	featureService      servicespec.FeatureService
	forwarderService    servicespec.ForwarderService
	fsService           servicespec.FSService
	idService           servicespec.IDService
	instrumentorService servicespec.InstrumentorService
	logService          servicespec.LogService
	networkService      servicespec.NetworkService
	permutationService  servicespec.PermutationService
	randomService       servicespec.RandomService
	storageCollection   servicespec.StorageCollection
	textInputService    servicespec.TextInputService
	textOutputService   servicespec.TextOutputService
	trackerService      servicespec.TrackerService

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Activator() servicespec.ActivatorService {
	return c.activatorService
}

func (c *collection) Boot() {
	go c.Activator().Boot()
	go c.Connection().Boot()
	go c.Endpoint().Boot()
	go c.Feature().Boot()
	go c.Forwarder().Boot()
	go c.FS().Boot()
	go c.ID().Boot()
	go c.Instrumentor().Boot()
	go c.Log().Boot()
	go c.Network().Boot()
	go c.Permutation().Boot()
	go c.Random().Boot()
	go c.Storage().Boot()
	go c.TextInput().Boot()
	go c.TextOutput().Boot()
	go c.Tracker().Boot()
}

func (c *collection) Connection() servicespec.ConnectionService {
	return c.connectionService
}

func (c *collection) Endpoint() servicespec.EndpointCollection {
	return c.endpointCollection
}

func (c *collection) Feature() servicespec.FeatureService {
	return c.featureService
}

func (c *collection) Forwarder() servicespec.ForwarderService {
	return c.forwarderService
}

func (c *collection) FS() servicespec.FSService {
	return c.fsService
}

func (c *collection) ID() servicespec.IDService {
	return c.idService
}

func (c *collection) Instrumentor() servicespec.InstrumentorService {
	return c.instrumentorService
}

func (c *collection) Log() servicespec.LogService {
	return c.logService
}

func (c *collection) Network() servicespec.NetworkService {
	return c.networkService
}

func (c *collection) Permutation() servicespec.PermutationService {
	return c.permutationService
}

func (c *collection) Random() servicespec.RandomService {
	return c.randomService
}

func (c *collection) SetActivatorService(activator servicespec.ActivatorService) {
	c.activatorService = activator
}

func (c *collection) SetConnectionService(connectionService servicespec.ConnectionService) {
	c.connectionService = connectionService
}

func (c *collection) SetEndpointCollection(endpointCollection servicespec.EndpointCollection) {
	c.endpointCollection = endpointCollection
}

func (c *collection) SetFeatureService(featureService servicespec.FeatureService) {
	c.featureService = featureService
}

func (c *collection) SetForwarderService(forwarderService servicespec.ForwarderService) {
	c.forwarderService = forwarderService
}

func (c *collection) SetFSService(fsService servicespec.FSService) {
	c.fsService = fsService
}

func (c *collection) SetIDService(idService servicespec.IDService) {
	c.idService = idService
}

func (c *collection) SetInstrumentorService(instrumentorService servicespec.InstrumentorService) {
	c.instrumentorService = instrumentorService
}

func (c *collection) SetLogService(logService servicespec.LogService) {
	c.logService = logService
}

func (c *collection) SetNetworkService(networkService servicespec.NetworkService) {
	c.networkService = networkService
}

func (c *collection) SetPermutationService(permutationService servicespec.PermutationService) {
	c.permutationService = permutationService
}

func (c *collection) SetRandomService(randomService servicespec.RandomService) {
	c.randomService = randomService
}

func (c *collection) SetStorageCollection(storageCollection servicespec.StorageCollection) {
	c.storageCollection = storageCollection
}

func (c *collection) SetTextInputService(textInputService servicespec.TextInputService) {
	c.textInputService = textInputService
}

func (c *collection) SetTextOutputService(textOutputService servicespec.TextOutputService) {
	c.textOutputService = textOutputService
}

func (c *collection) SetTrackerService(trackerService servicespec.TrackerService) {
	c.trackerService = trackerService
}

func (c *collection) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.Endpoint().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (c *collection) Storage() servicespec.StorageCollection {
	return c.storageCollection
}

func (c *collection) TextInput() servicespec.TextInputService {
	return c.textInputService
}

func (c *collection) TextOutput() servicespec.TextOutputService {
	return c.textOutputService
}

func (c *collection) Tracker() servicespec.TrackerService {
	return c.trackerService
}
