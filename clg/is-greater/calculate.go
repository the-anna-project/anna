// Package isgreater implements spec.CLG and provides a method to identify if
// the first given number is greater than the later.
package isgreater

// calculate checks whether the first given number is greater than the other.
func (c *clg) calculate(a, b float64) bool {
	return a > b
}
