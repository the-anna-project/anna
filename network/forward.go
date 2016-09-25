package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) forwardCLGs(ctx spec.Context, behaviorIDs []string, payload spec.NetworkPayload) error {
	for _, ID := range behaviorIDs {
		// Prepare a new context for the new connection path.
		newCtx := ctx.Clone()
		newCtx.SetBehaviorID(ID)

		// Create a new network payload. Note that the old context of the old
		// network payload is removed to only append actual arguments to the new
		// network payload.
		newPayloadConfig := api.DefaultNetworkPayloadConfig()
		newPayloadConfig.Args = append([]reflect.Value{reflect.ValueOf(newCtx)}, payload.GetArgs()[1:]...)
		newPayloadConfig.Destination = spec.ObjectID(ID)
		newPayloadConfig.Sources = []spec.ObjectID{payload.GetDestination()}
		newPayload, err := api.NewNetworkPayload(newPayloadConfig)
		if err != nil {
			return maskAny(err)
		}

		// Find the actual CLG based on its behavior ID. Therefore we lookup the
		// behavior name by the given behavior ID. Data we read here is written
		// within several CLGs. That way the network creates its own connections
		// based on learned experiences.
		//
		// TODO where are these connections coming from?
		// TODO if there are none, we need to find some randomly
		// TODO there needs to be some sort of variation when executing existing CLG trees
		//
		clgName, err := n.Storage().General().Get(key.NewCLGKey("behavior-id:%s:behavior-name", ID))
		if err != nil {
			return maskAny(err)
		}
		CLG, err := n.clgByName(clgName)
		if err != nil {
			return maskAny(err)
		}
		CLG.GetInputChannel() <- newPayload
	}

	return nil
}

// TODO
func (n *network) forwardInputCLG(payload spec.NetworkPayload) error {
	ctx, err := payload.GetContext()
	if err != nil {
		return maskAny(err)
	}

	// TODO fetch CLG tree ID using information ID
	// TODO fetch input CLG ID (behaviorID) using CLG tree ID
	//clgTreeID, err := c.Storage.Get(key.NewCLGKey("information-id:%s:clg-tree-id", informationID))
	//if storage.IsNotFound(err) {
	//	// We do not know any useful CLG tree for the given input. Thus we cannot
	//	// set any to the current context.
	//	return nil
	//} else if err != nil {
	//	return maskAny(err)

	var behaviorID string

	newCtx := ctx.Clone()
	newCtx.SetBehaviorID(behaviorID)

	// Create a new network payload
	newPayloadConfig := api.DefaultNetworkPayloadConfig()
	newPayloadConfig.Args = append([]reflect.Value{reflect.ValueOf(newCtx)}, payload.GetArgs()[1:]...)
	newPayloadConfig.Destination = spec.ObjectID(behaviorID)
	newPayloadConfig.Sources = []spec.ObjectID{payload.GetDestination()}
	newPayload, err := api.NewNetworkPayload(newPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	CLG, err := n.clgByName("input")
	if err != nil {
		return maskAny(err)
	}
	CLG.GetInputChannel() <- newPayload

	return nil
}

func (n *network) forwardOutputCLG(ctx spec.Context, payload spec.NetworkPayload) error {
	// Return the calculated output to the requesting client, if the
	// current CLG is the output CLG.
	newTextResponseConfig := api.DefaultTextResponseConfig()
	newTextResponseConfig.Output = payload.String()
	newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
	if err != nil {
		return maskAny(err)
	}

	n.TextOutput <- newTextResponse

	return nil
}

// helper

func (n *network) findConnections(ctx spec.Context, payload spec.NetworkPayload) ([]string, error) {
	var behaviorIDs []string

	behaviorID := ctx.GetBehaviorID()
	if behaviorID == "" {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}
	behaviorIDsKey := key.NewCLGKey("behavior-id:%s:behavior-ids", behaviorID)

	err := n.Storage().General().WalkSet(behaviorIDsKey, n.Closer, func(element string) error {
		behaviorIDs = append(behaviorIDs, element)
		return nil
	})
	if err != nil {
		return nil, maskAny(err)
	}

	return behaviorIDs, nil
}
