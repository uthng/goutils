package goutils_test

import (
	//"fmt"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uthng/goutils"
)

func TestStringConvertToMapSimple(t *testing.T) {
	newError := func(msg string) error {
		return errors.New("key/value element '" + msg + "' malformatted. Key and value must be separated by ':'")
	}
	testCases := []struct {
		name   string
		str    string
		result interface{}
	}{
		{
			"ErrKeySimple",
			"key1",
			newError("key1"),
		},
		{
			"ErrKeyWithoutValue",
			"key1:val1;key2",
			newError("key2"),
		},
		{
			"OKKVSimple",
			"key1:val1",
			map[string]interface{}{
				"key1": "val1",
			},
		},
		{
			"OKKVMultiple",
			"key1:val1;key2:val2",
			map[string]interface{}{
				"key1": "val1",
				"key2": "val2",
			},
		},
		{
			"ErrKVMultiple",
			"key1:val1;key2:val2;",
			newError(""),
		},
		{
			"OKKVMultipleNested",
			"key1:val1;key2:val2;key3:key31:val31",
			map[string]interface{}{
				"key1": "val1",
				"key2": "val2",
				"key3": map[string]interface{}{
					"key31": "val31",
				},
			},
		},
		{
			"ErrKVMultipleNested",
			"key1:val1;key2:val2;key3:key31:val31;key32",
			newError("key32"),
		},
		//{
		//"OKKVComplexeNested",
		//"key1:key11:val11,key12:val12;key2:val2;key3:key31:val31,key32:key321:val321,key322:val322;key33:val33",
		//map[string]interface{}{
		//"key1": map[string]interface{}{
		//"key11": "val11",
		//"key12": "val12",
		//},
		//"key2": "val2",
		//"key3": map[string]interface{}{
		//"key31": "val31",
		//"key32": map[string]interface{}{
		//"key321": "val321",
		//"key322": "val322",
		//},
		//"key33": "val33",
		//},
		//},
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := goutils.StringConvertToMapSimple(tc.str, ";", ":")
			if strings.HasPrefix(tc.name, "Err") {
				assert.Equal(t, tc.result, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.result, m)
		})
	}
}

func TestStringBuildWithSep(t *testing.T) {
	testCases := []struct {
		name   string
		strs   []string
		sep    rune
		result string
	}{
		{
			"OK",
			[]string{"a", "b", "c", "d"},
			';',
			"a;b;c;d",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := goutils.StringBuildWithSep(tc.sep, tc.strs...)
			require.Equal(t, tc.result, result)
		})
	}
}
