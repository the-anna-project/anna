package spec

// CLG represents the CLGs interacting with each other within the neural
// network. Each CLG is registered in the Network. From there signal are
// dispatched in a dynamic fashion until some useful calculation took place.
type CLG interface {
	FactoryProvider

	GatewayProvider

	// GetCalculate returns the CLG's calculate function which implements its
	// actual business logic.
	GetCalculate() interface{}

	// GetName returns the CLG's human readable name.
	GetName() string

	Object

	// SetFactoryCollection configures the CLG's factory collection. This is done
	// for all CLGs, regardless if a CLG is making use of the factory collection
	// or not.
	SetFactoryCollection(factoryCollection FactoryCollection)

	// SetGatewayCollection configures the CLG's gateway collection. This is done
	// for all CLGs, regardless if a CLG is making use of the gateway collection
	// or not.
	SetGatewayCollection(gatewayCollection GatewayCollection)

	// SetLog configures the CLG's logger. This is done for all CLGs, regardless
	// if a CLG is making use of the logger or not.
	SetLog(log Log)

	// SetStorageCollection configures the CLG's storage collection. This is done
	// for all CLGs, regardless if a CLG is making use of the storage collection
	// or not.
	SetStorageCollection(storageCollection StorageCollection)

	StorageProvider
}
