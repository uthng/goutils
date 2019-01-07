package goutils

import (
//    "fmt"
)

//type void* interface{}
type CompareFunc func(elem1 interface{}, elem2 interface{}) bool

// Check if an element is in an array of elements of the same type
//func IsElementInArray(elem interface{}, array []interface{}, fn CompareFunc) bool {
func IsElementInArray(elem interface{}, array []interface{}, fn CompareFunc) bool {

	for _, elm := range array {
		res := fn(elem, elm)
		if res == true {
			return true
		}
	}

	return false
}
