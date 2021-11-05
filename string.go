package goutils

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// StringStripAnsi removes ANSI escape code from string
func StringStripAnsi(str string) string {
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

	var re = regexp.MustCompile(ansi)

	return re.ReplaceAllString(str, "")
}

// StringConvertToMapSimple converts a string composed of key/value elements
// separated by a separator to a map. By default, the separator for key and value is ":"
// and the separator for each key/value element is ";"
func StringConvertToMapSimple(str, elSep, kvSep string) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	arrElems := strings.Split(str, elSep)
	for _, el := range arrElems {
		arrKVs := strings.SplitN(el, kvSep, 2)
		if len(arrKVs) != 2 {
			return nil, fmt.Errorf("key/value element '%s' malformatted. Key and value must be separated by '%s'", el, kvSep)
		}

		if strings.Contains(arrKVs[1], kvSep) {
			arr, err := StringConvertToMapSimple(arrKVs[1], elSep, kvSep)
			if err != nil {
				return nil, err
			}

			m[arrKVs[0]] = arr
		} else {
			m[arrKVs[0]] = arrKVs[1]
		}
	}

	return m, nil
}

// StringBuild concatenates strings together
func StringBuild(strs ...string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(str)
	}

	return sb.String()
}

// StringBuildWithSep concatenates strings together with a separator
func StringBuildWithSep(sep rune, strs ...string) string {
	var sb strings.Builder

	for idx, str := range strs {
		sb.Write([]byte(str))

		if idx < len(strs)-1 {
			sb.WriteRune(sep)
		}
	}

	return sb.String()
}

// StringParseFloat is an advance ParseFloat for golang, support scientific notation, comma separated number.
//
// Examples:
// StringParseFloat("1E2") = 100)
// StringParseFloat("1E-5") = 0.00001)
// StringParseFloat("1.6543E2") = 165.43)
// StringParseFloat("0.89E2") = 89)
// StringParseFloat("1.6543E-2") = 0.016543)
// StringParseFloat("156,819,129") = 156819129)
// StringParseFloat("156819129") = 156819129)
// StringParseFloat(".1E0") = 0.1)
// StringParseFloat(".1E1") = 1)
// StringParseFloat("0E1") = 0)
func StringParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}
