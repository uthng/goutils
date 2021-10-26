package goutils

import (
	"fmt"
	"regexp"
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
