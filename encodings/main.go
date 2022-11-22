package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
)

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
			fmt.Printf("failed to parse CSV. err=%v", err)
			return
		}

		reader := csv.NewReader(file)
		record, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("failed to parse CSV. err=%v", err)
			return
		}
		if name == "test_utf8_bom.csv" && record[0][0] != "\xef\xbb\xbf商品ID" {
			fmt.Println("failed to read utf8 with bom")
		}

		fmt.Printf("successfully parsed %s\n", name)
	}
}
