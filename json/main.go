package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	book := &Book{
		Title: "Once Upon A Time",
		Authors: []*Author{
			{
				Name: "John",
				Age:  32,
			},
			{
				Name: "Alice",
				Age:  28,
			},
		},
	}

	bookJSON, _ := json.Marshal(book)
	fmt.Printf("json encoded=%s\n", string(bookJSON))

	var b *Book
	json.Unmarshal(bookJSON, &b) // passing the pointer to b
	fmt.Printf("book=%v\n", b)
	fmt.Printf("authors at index 0=%+v\n", b.Authors[0])

}

// "json" tag defines the name of each field of encoded json string
type Book struct {
	Title   string    `json:"title"`
	Authors []*Author `json:"authors"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
