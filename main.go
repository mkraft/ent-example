package main

import (
	"context"
	"fmt"
	"log"

	"example.com/entexample/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Printf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

	user := client.User.Create().SetType("admin").SetName("Jane Doe").SaveX(context.Background())

	user, err = client.User.Get(context.Background(), user.ID)
	if err != nil {
		log.Printf("failed getting user: %v", err)
	}
	fmt.Printf("after create\n\t%v\n", user)

	user.Type = nil
	user.Name = nil
	fmt.Printf("before update\n\t%v\n", user)

	user.Update().SaveX(context.Background())
	user = client.User.GetX(context.Background(), user.ID)
	fmt.Printf("after update (fields are not persisted as null)\n\t%v\n", user)

	user.Update().ClearType().ClearName().SaveX(context.Background())
	user = client.User.GetX(context.Background(), user.ID)
	fmt.Printf("after update with ClearType and ClearName (fields are persisted as null)\n\t%v\n", user)
}
