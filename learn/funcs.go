package learn

// go have named return values
func NamedReturn() (res string) {
	res = "name res"
	return
}

// go can return multiple results
func MultipleReturn(x, y int) (sum, prod int) {
	return x + y, x * y
}
