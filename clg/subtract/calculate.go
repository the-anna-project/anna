// Package subtract implements spec.CLG and provides the mathematical operation
// of subtraction.
package subtract

// calculate creates the difference of the given float64s.
func (c *clg) calculate(a, b float64) float64 {
	return a - b
}
