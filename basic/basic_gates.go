package logicgates

func And3(a, b, c bool) bool {
	return And(And(a, b), c)
}
