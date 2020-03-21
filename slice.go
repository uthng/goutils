package goutils

import (
//    "fmt"
)

// CompareFunc is utility func to compare 2 elements
type CompareFunc func(elem1 interface{}, elem2 interface{}) bool

// SliceIsElemIn checks if an element is in an array of elements of the same type
func SliceIsElemIn(array []interface{}, elem interface{}, fn CompareFunc) int {

	for i, elm := range array {
		res := fn(elem, elm)
		if res == true {
			return i
		}
	}

	return -1
}

// SliceFindElemStr takes a slice and looks for an element in it. If found it will
// return the element's position. Otherwise it will return -1
func SliceFindElemStr(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
