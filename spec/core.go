package spec

type Core interface {
	// Boot initializes and starts the whole core like booting a machine. The
	// call to Boot blocks until the core is completely initialized, so you might
	// want to call it in a separate goroutine.
	Boot()

	GetObjectID() ObjectID

	GetObjectType() ObjectType

	GetState() State

	SetState(state State)

	// Shutdown ends all processes of the core like shutting down a machine. The
	// call to Boot blocks until the core is completely shut down, so you might
	// want to call it in a separate goroutine.
	Shutdown()

	Trigger(imp Impulse) (Impulse, error)
}
