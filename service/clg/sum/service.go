// Package sum implements spec.CLG and provides the mathematical operation of
// addition.
package sum

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// calculate creates the sum of the given float64s.
func (s *service) calculate(ctx spec.Context, a, b float64) float64 {
	return a + b
}
