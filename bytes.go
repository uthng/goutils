package goutils

import (
	"bytes"
	"encoding/gob"
)

// BytesMarshal converts a type interface to byte array
func BytesMarshal(i interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(i); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// BytesUnMarshal put data of a byte array to a type interface
func BytesUnMarshal(data []byte, i interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(i)
}

// BytesConvertToUTF8 converts bytes from a given encoding to UTF8
func BytesConvertToUTF8(bytes []byte, charset string) ([]byte, error) {
	encoding, err := getEncoding(charset)
	if err != nil {
		return nil, err
	}

	bs, err := encoding.NewDecoder().Bytes(bytes)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
