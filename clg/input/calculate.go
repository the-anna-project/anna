// Package input implements spec.CLG and provides the entry to the neural
// network.
package input

import (
	"fmt"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/storage"
)

// calculate tries to map the given input sequence to a CLG tree ID within the
// available storage.
func (c *clg) calculate(ctx context.Context, input string) error {
	informationIDKey := key.NewCLGKey(c, "input-sequence:information-id:%s", input)
	informationID, err := c.Storage.Get(informationIDKey)
	if storage.IsNotFound(err) {
		// The given input was never seen before. Thus we register it now with its
		// own very unique information ID.
		newID, err := c.IDFactory.WithType(id.Hex128)
		if err != nil {
			return maskAny(err)
		}
		err := c.Storage.Set(informationIDKey, newID)
		if err != nil {
			return maskAny(err)
		}

		// The given input is completely new, so we are not able to set a CLG tree
		// ID to the context. Thus we simply return here.
		return nil
	} else if err != nil {
		return maskAny(err)
	}
	clgTreeID, err := c.Storage.Get(key.NewCLGKey(c, "information-id:clg-tree-id:%s", informationID))
	if storage.IsNotFound(err) {
		// We do not know any useful CLG tree for the given input. Thus we cannot
		// set any to the current context.
		return nil
	} else if err != nil {
		return maskAny(err)
	}

	ctx.SetCLGTreeID(clgTreeID)

	return nil
}
