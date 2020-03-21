package goutils

import (
//    "fmt"
)

// CompareFunc is utility func to compare 2 elements
type CompareFunc func(elem1 interface{}, elem2 interface{}) bool

// SliceIsElemIn checks if an element is in an array of elements of the same type
func SliceIsElemIn(elem interface{}, array []interface{}, fn CompareFunc) bool {

	for _, elm := range array {
		res := fn(elem, elm)
		if res == true {
			return true
		}
	}

	return false
}

// SliceFindElemStr takes a slice and looks for an element in it. If found it will
// return the element's position and true. Otherwise it will return -1 and false
func SliceFindElemStr(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
