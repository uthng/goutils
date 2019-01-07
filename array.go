package goutils

import (
//    "fmt"
)

// CompareFunc is utility func to compare 2 elements
type CompareFunc func(elem1 interface{}, elem2 interface{}) bool

// ArrayIsElementIn checks if an element is in an array of elements of the same type
func ArrayIsElementIn(elem interface{}, array []interface{}, fn CompareFunc) bool {

	for _, elm := range array {
		res := fn(elem, elm)
		if res == true {
			return true
		}
	}

	return false
}
