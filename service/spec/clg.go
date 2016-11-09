package spec

import storagespec "github.com/xh3b4sd/anna/storage/spec"

// CLG represents the CLGs interacting with each other within the neural
// network. Each CLG is registered in the Network. From there signal are
// dispatched in a dynamic fashion until some useful calculation took place.
type CLG interface {
	Provider

	// GetCalculate returns the CLG's calculate function which implements its
	// actual business logic.
	GetCalculate() interface{}

	// GetMetadata returns the service's metadata.
	GetMetadata() map[string]string

	// SetServiceCollection configures the CLG's factory collection. This is done
	// for all CLGs, regardless if a CLG is making use of the factory collection
	// or not.
	SetServiceCollection(serviceCollection Collection)

	// SetStorageCollection configures the CLG's storage collection. This is done
	// for all CLGs, regardless if a CLG is making use of the storage collection
	// or not.
	SetStorageCollection(storageCollection storagespec.Collection)

	storagespec.Provider
}
