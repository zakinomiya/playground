package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
	"golang.org/x/text/transform"
)

type BOM []byte

var (
	UTF8    BOM = []byte{'\xEF', '\xBB', '\xBF'}
	UTF16LE     = []byte{'\xFF', '\xFE'}
	UTF16BE     = []byte{'\xFE', '\xFF'}
	UTF32LE     = []byte{'\xFF', '\xFE', '\x00', '\x00'}
	UTF32BE     = []byte{'\x00', '\x00', '\xFE', '\xFF'}
)

// bomLength returns the length of the BOM in the given file
// If the file does not include the BOM, it will return 0
func bomLength(b []byte) int64 {
	fmt.Printf("%x\n", b)
	for _, bom := range []BOM{UTF8, UTF16LE, UTF16BE, UTF32LE, UTF32BE} {
		if bytes.HasPrefix(b, bom) {
			fmt.Println(bom)
			return int64(len(bom))
		}
	}
	return 0
}

func trimBom(file *os.File) (*os.File, error) {
	b := make([]byte, 4)
	_, err := file.Read(b)
	if err != nil {
		return nil, err
	}

	// skip the BOM if contained
	if l := bomLength(b); l != 0 {
		file.Seek(l, 0)
	} else {
		file.Seek(0, 0)
	}

	return file, nil
}

func getDecoder(b []byte) *encoding.Decoder {
	if bytes.HasPrefix(b, UTF8) {
		fmt.Println("utf8 bom")
		return unicode.UTF8BOM.NewDecoder()
	} else if bytes.HasPrefix(b, UTF16LE) {
		fmt.Println("utf16 le")
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	} else if bytes.HasPrefix(b, UTF16BE) {
		fmt.Println("utf16 be")
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
	} else if bytes.HasPrefix(b, UTF32LE) {
		fmt.Println("utf32 le")
		return utf32.UTF32(utf32.BigEndian, utf32.UseBOM).NewDecoder()
	} else if bytes.HasPrefix(b, UTF32BE) {
		fmt.Println("utf32 be")
		return utf32.UTF32(utf32.BigEndian, utf32.UseBOM).NewDecoder()
	}

	// fmt.Println("utf8")
	return unicode.UTF8.NewDecoder()
}

func GeneralCSVReader(file *os.File) (*csv.Reader, error) {
	b := make([]byte, 4)
	_, err := file.Read(b)
	if err != nil {
		return nil, err
	}
	file.Seek(0, 0)

	return csv.NewReader(transform.NewReader(file, getDecoder(b))), nil
}

func main() {
	files, err := os.ReadDir("csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		name := f.Name()
		file, err := os.OpenFile(path.Join("csv", name), os.O_RDONLY, 0644)
		if err != nil {
			fmt.Printf("file=%s: failed to open CSV. err=%v", name, err)
			continue
		}

		// file, err = trimBom(file)
		// if err != nil {
		// 	fmt.Printf("failed to trim the BOM. err=%v", err)
		// 	return
		// }

		// reader := csv.NewReader(file)
		reader, err := GeneralCSVReader(file)
		if err != nil {
			fmt.Printf("file=%s: failed to get GeneralCSVReader. err=%v\n", name, err)
			continue
		}
		_, err = reader.ReadAll()
		if err != nil {
      fmt.Printf("err: file=%s: failed to parse CSV. err=%v\n", name, err)
			continue
		}

		// fmt.Printf("successfully parsed %s\n", name)
	}
}
