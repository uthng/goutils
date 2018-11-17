package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMapSortedKeys(t *testing.T) {
	map1 := map[string]int{
		"bj": 3,
		"ba": 4,
		"ac": 6,
		"cc": 7,
	}

	res1 := []string{"ac", "ba", "bj", "cc"}
	res2 := []string{"cc", "bj", "ba", "ac"}

	sortedKeys := GetMapSortedKeys(map1, false)
	fmt.Println(sortedKeys)
	assert.Equal(t, res1, sortedKeys)

	sortedKeys = GetMapSortedKeys(map1, true)
	fmt.Println(sortedKeys)
	assert.Equal(t, res2, sortedKeys)

}

func TestGetMapKeys(t *testing.T) {
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
			o, err := GetMapKeys(tc.input)
			assert.Nil(t, err)

			assert.ElementsMatch(t, o, tc.output)
		})
	}
}
