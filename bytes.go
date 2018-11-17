package utils

import (
	"bytes"
	"encoding/gob"
)

// MarshalBytes converts a type interface to byte array
func MarshalBytes(i interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(i); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnMarshalBytes put data of a byte array to a type interface
func UnMarshalBytes(data []byte, i interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(i)
}
