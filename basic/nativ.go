// +build nativ

// Logicgates
package logicgates

// Not gate
func Not(a bool) bool {
	return !a
}

// And gate
// Returns 1 if both a and b is, 0 otherwise.
func And(a, b bool) bool {
	return (a && b)
}

// Or gate
// Returns 0 if both inputs is 0, 1 otherwise.
func Or(a, b bool) bool {
	return (a || b)
}

// Nor function
// Returns 1 if both a and b is 0, and returns 0 in every other cases.
func Nor(a, b bool) bool {
	return !(a || b)
}

// Nand function
// Returns 0 if both a and b is 1, and returns 1 in every other cases.
func Nand(a, b bool) bool {
	return !(a && b)
}
