package utils

import (
	"fmt"
	"reflect"
	"sort"
)

// GetMapSortedKeys returns a sorted slice contaning keys of map as interface{}// and need to be casted in the format as you want
//
// This function supports the following type of key: string, int
// If the given parameter is not a map, it will return nil
func GetMapSortedKeys(m interface{}, reverse bool) interface{} {
	var sortedInts []int
	var sortedStrings []string
	//var sortedKeys reflect.Value

	mType := reflect.TypeOf(m)
	mValue := reflect.ValueOf(m)

	if mType.Kind() != reflect.Map {
		return nil
	}

	if mType.Key().Kind() == reflect.String {
		//sortedKeys = reflect.MakeSlice(reflect.TypeOf(sstring), 0, mValue.Len())
		for _, val := range mValue.MapKeys() {
			sortedStrings = append(sortedStrings, val.Interface().(string))
		}
		if reverse == false {
			sort.Strings(sortedStrings)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(sortedStrings)))
		}

		return sortedStrings
	} else if mType.Key().Kind() == reflect.Int {
		//sortedKeys = reflect.MakeSlice(reflect.TypeOf(sint), 0, mValue.Len())
		for _, val := range mValue.MapKeys() {
			sortedInts = append(sortedInts, int(val.Int()))
		}
		if reverse == false {
			sort.Ints(sortedInts)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(sortedInts)))
		}

		return sortedInts
	}

	return nil
}

// GetMapKeys returns a slice of keys.
func GetMapKeys(m interface{}) ([]interface{}, error) {
	var out []interface{}

	v := reflect.ValueOf(m)

	if v.Kind() != reflect.Map {
		return out, fmt.Errorf("Input argument is not a map")
	}

	for _, val := range v.MapKeys() {
		out = append(out, val.Interface())
	}

	return out, nil
}
