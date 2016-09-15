// Package splitfeatures implements spec.CLG and provides functionality to
// split information sequences into features.
package splitfeatures

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/index/clg/collection/feature-set"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

const (
	FeatureSize int = 4
)

func (c *clg) calculate(ctx spec.Context, informationSequence string) error {
	newConfig := featureset.DefaultConfig()
	newConfig.MaxLength = FeatureSize
	newConfig.MinLength = FeatureSize
	newConfig.Sequences = []string{informationSequence}
	newFeatureSet, err := featureset.New(newConfig)
	if err != nil {
		return maskAny(err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		return maskAny(err)
	}

	features := newFeatureSet.GetFeatures()
	for _, f := range features {
		positionKey := key.NewCLGKey("feature:%s:positions", f.GetSequence())
		raw, err := json.Marshal(f.GetPositions())
		if err != nil {
			return maskAny(err)
		}
		err = c.Storage.Set(positionKey, string(raw))
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
