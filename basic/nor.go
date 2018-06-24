// +build nor

// Universal gate based on nor gate
package logicgates

// Nor function
// Returns 1 if both a and b is 0, and returns 0 in every other cases.
func Nor(a, b bool) bool {
	return !(a || b)
}

// Build other gates from Nor and Nor alone
// Not gate
// Nor(1,1) = 0; Nor(0,0) = 1
func Not(a bool) bool {
	return Nor(a, a)
}

// Build other gates from Nor and Nor alone
// And gate
// Returns 1 if both a and b is, 0 otherwise.
func And(a, b bool) bool {
	return Nor(Nor(a, a), Nor(b, b))
}

// Build other gates from Nor and Nor alone
// Or gate
// Returns 0 if both inputs is 0, 1 otherwise.
func Or(a, b bool) bool {
	return Nor(Nor(a, b), Nor(a, b))
}

// Nand function
// Returns 0 if both a and b is 1, and returns 1 in every other cases.
func Nand(a, b bool) bool {
	return Nor(Nor(Nor(a, a), Nor(b, b)), Nor(Nor(a, a), Nor(b, b)))
}
