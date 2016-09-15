// Package informationid implements spec.CLG and provides functionality to read
// the information sequence stored under a specific information ID.
package readinformationid

import (
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

// calculate fetches the information sequence stored under a specific
// information ID. The information ID is provided by the given context.
func (c *clg) calculate(ctx spec.Context) (string, error) {
	informationID := ctx.GetInformationID()
	if informationID == "" {
		return "", maskAnyf(informationIDError, "must not be empty")
	}

	informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", informationID)
	informationSequence, err = c.Storage.Get(informationSequenceKey)
	if err != nil {
		return "", maskAny(err)
	}

	return informationSequence, nil
}