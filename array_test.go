package goutils

import (
	"fmt"
	"testing"
)

func TestStrInArrayStr(t *testing.T) {
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

	res = IsElementInArray(str1, arrayStr, fn)
	if res == true {
		fmt.Printf("OK: \"%v\" is found in [%v]\n", str1, arrayStr)
		testRes = true
	} else {
		fmt.Printf("ERR: \"%v\" is not found in [%v]\n", str1, arrayStr)
	}

	if testRes == false {
		t.Fail()
	}

	res = IsElementInArray(str2, arrayStr, fn)
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
