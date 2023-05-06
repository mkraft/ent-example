package main

import (
	"context"
	"log"

	"example.com/entexample/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	user := client.User.Create().SetType("admin").SaveX(context.Background())

	user, err = client.User.Get(context.Background(), user.ID)
	if err != nil {
		log.Fatalf("failed getting user: %v", err)
	}
	log.Println("after save\t\t", user)

	user.Type = nil
	log.Println("before update\t", user)

	user = user.Update().SaveX(context.Background())
	log.Println("after update\t", user)

	user = user.Update().ClearType().SaveX(context.Background())
	log.Println("after cleartype\t", user)
}
