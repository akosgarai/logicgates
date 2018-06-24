package logicgates

// And3 - And gate with 3 inputs
func And3(a, b, c bool) bool {
	return And(And(a, b), c)
}

// Xor gate
func Xor(a, b bool) bool {
	return Or(And(a, Not(b)), And(b, Not(a)))
}

// Xnor gate
func Xnor(a, b bool) bool {
	return Not(Xor(a, b))
}
