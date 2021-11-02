package goutils

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

// MapGetSortedKeys returns a sorted slice contaning keys of map as interface{}// and need to be casted in the format as you want
//
// This function supports the following type of key: string, int
// If the given parameter is not a map, it will return nil
func MapGetSortedKeys(m interface{}, reverse bool) interface{} {
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

		if !reverse {
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

		if !reverse {
			sort.Ints(sortedInts)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(sortedInts)))
		}

		return sortedInts
	}

	return nil
}

// MapGetKeys returns a slice of keys.
func MapGetKeys(m interface{}) ([]interface{}, error) {
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

// MapStringMerge merges 2 string maps of the same type: map2 in map1.
//
// If an element in map1 does not exist, it will be created.
// Otherwise, it will be overrided by new value in map2
func MapStringMerge(map1, map2 interface{}) (interface{}, error) {
	v1 := reflect.ValueOf(map1)
	v2 := reflect.ValueOf(map2)

	t1 := v1.Type()
	t2 := v2.Type()

	if !strings.HasPrefix(t1.String(), "map[string]") || !strings.HasPrefix(t2.String(), "map[string]") {
		return nil, fmt.Errorf("one of 2 two maps is not string map")
	}

	if t1 != t2 {
		return nil, fmt.Errorf("two map are not the same type of string map")
	}

	m1 := cast.ToStringMap(map1)
	m2 := cast.ToStringMap(map2)

	m := m1

	for k, v := range m2 {
		m[k] = v
	}

	return m, nil
}

// MapRemoveNulls removes all fields with nil value
func MapRemoveNulls(m map[string]interface{}) {
	val := reflect.ValueOf(m)

	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}

		switch t := v.Interface().(type) {
		// If key is a JSON object (Go Map), use recursion to go deeper
		case map[string]interface{}:
			MapRemoveNulls(t)
		}
	}
}
