package learn

import (
	"fmt"
)

func PrimitiveTypes() string {
	str := "go"                                      // short declaration
	var str2 = "has different variable declarations" // var syntax
	return str + " " + str2
}

// define point as a struct with two fields (int x and y)
type point struct {
	x, y int
}

// define stringer interface with one method: String()
type Stringer interface {
	String() string
}

// define a method on type pair
// implements Stringer interface automatically
func (p point) String() string {
	return fmt.Sprintf("point %d, %d", p.x, p.y)
}

func Interfaces() Stringer {
	p := point{5, 6} // struct literal brace syntax

	var i Stringer
	i = p // pair implements Stringer

	return i // println will call String() method
}
