package goutils_test

import (
	//"fmt"
	//"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uthng/goutils"
)

func TestSliceIsElemIn(t *testing.T) {

	testCases := []struct {
		name   string
		elem   string
		result interface{}
	}{
		{
			"ErrElemNotFound",
			"mimi",
			-1,
		},
		{
			"OKElemFound",
			"tata",
			2,
		},
	}

	sliceStr := []interface{}{"toto", "titi", "tata", "tete"}

	fn := func(s1 interface{}, s2 interface{}) bool {

		if s1 == s2 {
			return true
		}
		return false
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pos := goutils.SliceIsElemIn(sliceStr, tc.elem, fn)
			assert.Equal(t, tc.result, pos)
		})
	}
}

func TestSliceFindElemStr(t *testing.T) {
	testCases := []struct {
		name   string
		elem   string
		result interface{}
	}{
		{
			"ErrElemNotFound",
			"mimi",
			-1,
		},
		{
			"OKElemFound",
			"tata",
			2,
		},
	}

	sliceStr := []string{"toto", "titi", "tata", "tete"}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pos := goutils.SliceFindElemStr(sliceStr, tc.elem)
			assert.Equal(t, tc.result, pos)
		})
	}
}
