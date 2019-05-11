// +build nand

// Universal gate based on nand gate
package logicgates

// Nand function
// Returns 0 if both a and b is 1, and returns 1 in every other cases.
func Nand(a, b bool) bool {
	return !(a && b)
}

// Build other gates from Nand and Nand alone
// Not gate
// Nand(1,1) = 0; Nand(0,0) = 1
func Not(a bool) bool {
	return Nand(a, a)
}

// Build other gates from Nand and Nand alone
// And gate
// Returns 1 if both a and b is, 0 otherwise.
func And(a, b bool) bool {
	return Nand(Nand(a, b), Nand(a, b))
}

// Build other gates from Nand and Nand alone
// Or gate
// Returns 0 if both inputs is 0, 1 otherwise.
func Or(a, b bool) bool {
	return Nand(Nand(a, a), Nand(b, b))
}

// Nor function
// Returns 1 if both a and b is 0, and returns 0 in every other cases.
func Nor(a, b bool) bool {
	return Nand(Nand(Nand(a, a), Nand(b, b)), Nand(Nand(a, a), Nand(b, b)))
}
