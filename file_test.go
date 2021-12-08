package goutils_test

import (
	//"fmt"
	//"strings"

	"fmt"
	"io"
	"os"
	"testing"

	//"github.com/endeveit/enca"

	"github.com/stretchr/testify/require"
	"github.com/uthng/goutils"
)

func TestFileGuessEncoding(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		result   string
	}{
		{
			"iso-8859-1",
			"iso8859-1.txt",
			`Prénom;Nom;Société;email;;;
abcàâäéèêë;îôöùûüabç;Money?$£;nom.prenom+testCSV@toto.co;;;`,
		},
		//{
		//"ascii",
		//"ascii.txt",
		//true,
		//},
		////{
		////"cp-865",
		////"cp865.txt",
		////true,
		////},
		//{
		//"koi8-r",
		//"koi8_r.txt",
		//true,
		//},
		//{
		//"latin1",
		//"latin1.txt",
		//true,
		//},
		{
			"utf-16le",
			"utf-16le.txt",
			"\ufeffpremiÈre is first\npremière is slightly different\nКириллица is Cyrillic\n𐐀 am Deseret",
		},
		{
			"utf-8",
			"utf8.txt",
			"10€ est chère",
		},
		{
			"utf-8bom",
			"utf8-bom.csv",
			`Prénom;Nom;email;Téléphone;Intérêts
Luçile;Rivière;lucile.riviere+pc@plezi.co;+33145048955;cheval`,
		},
		//{
		//"iso-8859-15",
		//"iso8859-15.txt",
		//true,
		//},
		{
			"windows-1252",
			"windows-1252.csv",
			`Prénom;Nom;Société;email;;;
abcàâäéèêë;îôöùûüabç;Money€$£;lucile.riviere+testCSV@plezi.co;;;`,
		},
		//{
		//"cp865",
		//"QA_import_CSV_manyweirdcarac_MSDOS_pv.csv",
		//true,
		//},
		//{
		//"x-mac-romain",
		//"QA_import_CSV_manyweirdcarac_classic_pv.csv",
		//true,
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open("fixtures/" + tc.filename)
			require.Nil(t, err)

			b, err := io.ReadAll(f)
			require.Nil(t, err)

			//fmt.Printf("hex: % x\n", b)
			//fmt.Printf("uni: %#U\n", b)
			enc := goutils.FileGuessEncoding(b)

			fmt.Println("enc", enc)

			r, err := goutils.BytesConvertToUTF8(b, enc)
			require.Nil(t, err)

			require.Equal(t, tc.result, string(r))

			require.Equal(t, tc.name, enc)
		})
	}
}
