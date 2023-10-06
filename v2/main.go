// Package gotrain is a training package that helps teach how to set up GO packages.
//
// It includes one function Add that adds two integers together.
package gotrain

import "golang.org/x/exp/constraints"

// Number represents either an integer or a floating point number.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add preforms [additon] on two integers and returns the result.
//
// [additon]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x1, x2 T) T {
	return x1 + x2
}