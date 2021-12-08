package goutils

import (
	"fmt"
	"io"
	"regexp"
)

// FileConvertToUTF8 converts a file from a given encoding to UTF-8
func FileConvertToUTF8(file io.Reader, charset string) (io.Reader, error) {
	encoding, err := getEncoding(charset)
	if err != nil {
		return nil, err
	}

	return encoding.NewDecoder().Reader(file), nil
}

// FileGuessEncoding tries to guess the encoding of bytes.
//
// It recognizes UTF-8, UTF-32LE, UTF-32BE, UTF-16LE, UTF-16BE, windows-1252, iso-8859-1.
func FileGuessEncoding(bytes []byte) string {
	strQ := fmt.Sprintf("%+q", bytes)

	// Clean double quote at the begining & at the end
	if strQ[0] == '"' {
		strQ = strQ[1:]
	}

	if strQ[len(strQ)-1] == '"' {
		strQ = strQ[0 : len(strQ)-1]
	}

	// If utf-8-bom, it must start with \ufeff
	re := regexp.MustCompile(`^\\ufeff`)

	found := re.MatchString(strQ)
	if found {
		return "utf-8bom"
	}

	// If utf-8, it must contain \uxxxx
	re = regexp.MustCompile(`\\u[a-z0-9]{4}`)

	found = re.MatchString(strQ)
	if found {
		return "utf-8"
	}

	// utf-32be
	re = regexp.MustCompile(`^\\x00\\x00\\xfe\\xff`)

	found = re.MatchString(strQ)
	if found {
		return "utf-32be"
	}

	// utf-32le
	re = regexp.MustCompile(`^\\xff\\xfe\\x00\\x00`)

	found = re.MatchString(strQ)
	if found {
		return "utf-32le"
	}

	// utf-16be
	re = regexp.MustCompile(`^\\xff\\xff`)

	found = re.MatchString(strQ)
	if found {
		return "utf-16be"
	}

	// utf-16le
	re = regexp.MustCompile(`^\\xff\\xfe`)

	found = re.MatchString(strQ)
	if found {
		return "utf-16le"
	}

	// Check if 0x8{0-F} or 0x9{0-F} is present
	re = regexp.MustCompile(`(\\x8[0-9a-f]{1}|\\x9[0-9a-f]{1})`)

	found = re.MatchString(strQ)
	if found {
		// It might be windows-1252 or mac-roman
		// But at this moment, do not have mean to distinguish both. So fallback to windows-1252
		return "windows-1252"
	}

	// No 0x8{0-F} or 0x9{0-F} found, it might be iso-8859-xx
	// We tried to detect whether it is iso-8859-1 or iso-8859-15
	// Check if 0x8{0-F} or 0x9{0-F} is present
	//re = regexp.MustCompile(`(\\xa[4|6|8]{1}|\\xb[4|8|c|d|e]{1})`)

	//loc := re.FindStringIndex(strQ)
	//if loc != nil {
	//c := strQ[loc[0]:loc[1]]
	//fmt.Printf("char %s\n", c)
	//if enc, err := BytesConvertToUTF8(bytes, "iso-8859-15"); err == nil {
	//fmt.Println("converted bytes", enc)
	//fmt.Printf("converted %+q\n", enc)
	//}

	// At this moment, we can not detect the difference between iso-8859-x.
	// So just return a fallback iso-8859-1
	return "iso-8859-1"
}

//var charISO88591 = map[string]struct {
//r string
//b []byte
//}{
//"\\xa4": {"¤", []byte{0xa4}},
//"\\xa6": {"¦", []byte{0xa6}},
//"\\xa8": {"¨", []byte{0xa8}},
//"\\xb4": {"´", []byte{0xb4}},
//"\\xb8": {"¸", []byte{0xb8}},
//"\\xbc": {"¼", []byte{0xbc}},
//"\\xbd": {"½", []byte{0xbd}},
//"\\xbe": {"¾", []byte{0xbe}},
//}

//var charISO885915 = map[string]struct {
//r string
//b []byte
//}{
//"\\xa4": {"€", []byte{0xa4}},
//"\\xa6": {"Š", []byte{0xa6}},
//"\\xa8": {"š", []byte{0xa8}},
//"\\xb4": {"Ž", []byte{0xb4}},
//"\\xb8": {"ž", []byte{0xb8}},
//"\\xbc": {"Œ", []byte{0xbc}},
//"\\xbd": {"œ", []byte{0xbd}},
//"\\xbe": {"Ÿ", []byte{0xbe}},
//}
