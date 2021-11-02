package goutils

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapGetSortedKeys(t *testing.T) {
	map1 := map[string]int{
		"bj": 3,
		"ba": 4,
		"ac": 6,
		"cc": 7,
	}

	res1 := []string{"ac", "ba", "bj", "cc"}
	res2 := []string{"cc", "bj", "ba", "ac"}

	sortedKeys := MapGetSortedKeys(map1, false)
	assert.Equal(t, res1, sortedKeys)

	sortedKeys = MapGetSortedKeys(map1, true)
	assert.Equal(t, res2, sortedKeys)
}

func TestMapGetKeys(t *testing.T) {
	testCases := []struct {
		name   string
		input  interface{}
		output interface{}
	}{
		{
			"map[string]interface{}",
			map[string]interface{}{
				"key1": "val1",
				"key2": 2,
				"key3": []int{1, 2, 3},
				"key4": map[string]interface{}{
					"key41": "val41",
					"key42": map[string]interface{}{
						"key421": "val421",
						"key422": []string{"one", "two", "three"},
					},
				},
			},
			[]interface{}{"key1", "key2", "key3", "key4"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			o, err := MapGetKeys(tc.input)
			assert.Nil(t, err)

			assert.ElementsMatch(t, o, tc.output)
		})
	}
}

func TestMapStringMerge(t *testing.T) {
	testCases := []struct {
		name   string
		arg1   interface{}
		arg2   interface{}
		output interface{}
	}{
		{
			"NotMapString",
			map[int]string{
				1: "value",
			},
			map[string]string{
				"key": "value",
			},
			"one of 2 two maps is not string map",
		},
		{
			"NotSameTypeMapString",
			map[string]string{
				"key": "value",
			},
			map[string]bool{
				"key": true,
			},
			"two map are not the same type of string map",
		},
		{
			"MapStringInterface",
			map[string]interface{}{
				"key1": 1,
				"key2": true,
				"key3": "value3",
			},
			map[string]interface{}{
				"key4": true,
				"key5": 5,
				"key2": "value2",
				"key3": []int{1, 2, 3},
			},
			map[string]interface{}{
				"key1": 1,
				"key2": "value2",
				"key3": []int{1, 2, 3},
				"key4": true,
				"key5": 5,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := MapStringMerge(tc.arg1, tc.arg2)
			if err != nil {
				assert.Equal(t, tc.output, err.Error())
			} else {
				assert.Equal(t, tc.output, res)
			}
		})
	}
}

func TestMapRemoveNulls(t *testing.T) {
	testCases := []struct {
		name   string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			"OKRemoveNulls",
			map[string]interface{}{
				"key1": nil,
				"key2": 2,
				"key3": []int{1, 2, 3},
				"key4": map[string]interface{}{
					"key41": nil,
					"key42": map[string]interface{}{
						"key421": "val421",
						"key422": nil,
					},
				},
			},
			map[string]interface{}{
				"key2": 2,
				"key3": []int{1, 2, 3},
				"key4": map[string]interface{}{
					"key42": map[string]interface{}{
						"key421": "val421",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := tc.input

			MapRemoveNulls(m)

			assert.Equal(t, tc.output, m)
		})
	}
}
