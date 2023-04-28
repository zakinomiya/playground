package main

import (
	"context"
	"log"
)

func main() {
	db := newDB()
	ctx := context.Background()

	tables := []Model{User{}}
	if err := db.CreateTables(ctx, tables); err != nil {
		// log.Fatalln(err)
	}

	for i := 0; i < 10; i++ {
		if err := newRandomUser().save(ctx, db); err != nil {
			log.Fatalln(err)
		}
		log.Println(i)
	}
}
