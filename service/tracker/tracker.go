package tracker

import (
	"sync"

	"github.com/xh3b4sd/anna/key"
	objectspec "github.com/xh3b4sd/anna/object/spec"
	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/id"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

// Config represents the configuration used to create a new tracker object.
type Config struct {
	// Dependencies.
	ServiceCollection servicespec.Collection
	StorageCollection storagespec.Collection
}

// DefaultConfig provides a default configuration to create a new tracker object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		ServiceCollection: service.MustNewCollection(),
		StorageCollection: storage.MustNewCollection(),
	}

	return newConfig
}

// New creates a new configured tracker object.
func New(config Config) (systemspec.Tracker, error) {
	newTracker := &tracker{
		Config: config,

		Metadata: map[string]string{
			"id":   id.MustNewID(),
			"name": "tracker",
			"type": "service",
		},
	}

	if newTracker.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newTracker.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	return newTracker, nil
}

// MustNew creates either a new default configured tracker object, or panics.
func MustNew() systemspec.Tracker {
	newTracker, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newTracker
}

type tracker struct {
	Config

	Metadata map[string]string
}

func (t *tracker) CLGIDs(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error {
	destinationID := string(networkPayload.GetDestination())
	sourceIDs := networkPayload.GetSources()

	errors := make(chan error, len(sourceIDs))
	wg := sync.WaitGroup{}

	for _, s := range sourceIDs {
		wg.Add(1)
		go func(s string) {
			// Persist the single CLG ID connections.
			behaviourIDKey := key.NewNetworkKey("behaviour-id:%s:o:tracker:behaviour-ids", s)
			err := t.Storage().General().PushToSet(behaviourIDKey, destinationID)
			if err != nil {
				errors <- maskAny(err)
			}
			wg.Done()
		}(string(s))
	}

	wg.Wait()

	select {
	case err := <-errors:
		if err != nil {
			return maskAny(err)
		}
	default:
		// Nothing do here. No error occurred. All good.
	}

	return nil
}

func (t *tracker) CLGNames(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error {
	destinationName := CLG.GetName()
	sourceIDs := networkPayload.GetSources()

	errors := make(chan error, len(sourceIDs))
	wg := sync.WaitGroup{}

	for _, s := range sourceIDs {
		wg.Add(1)
		go func(s string) {
			behaviourNameKey := key.NewNetworkKey("behaviour-id:%s:behaviour-name", s)
			name, err := t.Storage().General().Get(behaviourNameKey)
			if err != nil {
				errors <- maskAny(err)
			} else {
				// The errors channel is capable of buffering one error for each source
				// ID. The else clause is necessary to queue only one possible error for
				// each source ID. So in case the name lookup was successful, we are
				// able to actually persist the single CLG name connection.
				behaviourNameKey := key.NewNetworkKey("behaviour-name:%s:o:tracker:behaviour-names", name)
				err := t.Storage().General().PushToSet(behaviourNameKey, destinationName)
				if err != nil {
					errors <- maskAny(err)
				}
			}

			wg.Done()
		}(string(s))
	}

	wg.Wait()

	select {
	case err := <-errors:
		if err != nil {
			return maskAny(err)
		}
	default:
		// Nothing do here. No error occurred. All good.
	}

	return nil
}

func (t *tracker) Track(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error {
	t.Service().Log().Line("func", "Track")

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error{
		t.CLGIDs,
		t.CLGNames,
	}

	// Execute one lookup after another to track connection path patterns.
	var err error
	for _, l := range lookups {
		err = l(CLG, networkPayload)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
