package spec

// Activator represents an management layer to organize CLG activation rules.
// The activator obtains network payloads for every single requested CLG of
// every possible CLG tree.
type Activator interface {
	// Activate represents the public interface that bundles the following lookup
	// functions.
	//
	//     GetNetworkPayload
	//     NewNetworkPayload
	//
	// Activate fetches the list of all queued network payloads of the requested
	// CLG from the underlying storage. The stored list is merged with the given
	// network payload and provided to the lookup functions listed above. Once
	// Activate found a matching network payload, the network payloads it is made
	// of are removed from the stored queue and the created network payload is
	// returned. The modifications of the updated queue are also persisted.
	Activate(CLG CLG, networkPayload NetworkPayload) (NetworkPayload, error)

	FactoryProvider

	// GetNetworkPayload compares the given queue against the stored configuration
	// of the requested CLG. This configuration is a combination of behavior IDs
	// that are known to be successful. We know that this configuration was
	// already successful in the past when it was created by newNetworkPayload
	// beforehand. Such a creation then happened in some CLG tree execution before
	// the current one. Anyway, in case the queue of network payloads given
	// contains network payloads sent by the CLGs listed in the stored
	// configuration, the interface of the requested CLG is fulfilled. Then a new
	// network payload is created by merging the matching network payloads of the
	// stored queue. In case no activation configuration of the requested CLG is
	// stored, or no match using the stored configuration associated with the
	// requested CLG can be found, an error is returned.
	GetNetworkPayload(CLG CLG, queue []NetworkPayload) (NetworkPayload, error)

	// NewNetworkPayload uses the given queue to find a combination of network
	// payloads that fulfill the interface of the requested CLG. This creation
	// process may be random or biased in some way. In case some created
	// combination of network payloads fulfills the interface of the requested
	// CLG, this found combination is stored as activation configuration for the
	// requested CLG. In case no match using the permuted network payloads of the
	// given queue can be found, an error is returned.
	NewNetworkPayload(CLG CLG, queue []NetworkPayload) (NetworkPayload, error)

	StorageProvider
}
