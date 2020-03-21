package goutils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uthng/goutils"
)

func TestSliceIsElemIn(t *testing.T) {
	var res bool
	arrayStr := []interface{}{"toto", "titi", "tata", "tete"}
	str1 := "titi"
	str2 := "blabla"
	testRes := false

	fn := func(s1 interface{}, s2 interface{}) bool {

		if s1 == s2 {
			return true
		}
		return false
	}

	res = goutils.SliceIsElemIn(str1, arrayStr, fn)
	if res == true {
		fmt.Printf("OK: \"%v\" is found in [%v]\n", str1, arrayStr)
		testRes = true
	} else {
		fmt.Printf("ERR: \"%v\" is not found in [%v]\n", str1, arrayStr)
	}

	if testRes == false {
		t.Fail()
	}

	res = goutils.SliceIsElemIn(str2, arrayStr, fn)
	if res == true {
		fmt.Printf("ERR: \"%v\" is found in [%v]\n", str2, arrayStr)
	} else {
		fmt.Printf("OK: \"%v\" is not found in [%v]\n", str2, arrayStr)
		testRes = true
	}

	if testRes == false {
		t.Fail()
	}

}

func TestSliceFindElemStr(t *testing.T) {
	testCases := []struct {
		name   string
		slice  []string
		elem   string
		result interface{}
	}{
		{
			"ErrElemNotFound",
			[]string{"toto", "titi", "tata", "tete"},
			"mimi",
			-1,
		},
		{
			"OKElemFound",
			[]string{"toto", "titi", "tata", "tete"},
			"tata",
			2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pos, ok := goutils.SliceFindElemStr(tc.slice, tc.elem)
			if strings.HasPrefix(tc.name, "Err") {
				assert.Equal(t, ok, false)
				assert.Equal(t, pos, tc.result)
				return
			}

			assert.Equal(t, ok, true)
			assert.Equal(t, pos, tc.result)
		})
	}
}
