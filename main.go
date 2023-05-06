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
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

	user := client.User.Create().SetType("admin").SaveX(context.Background())

	_, err = client.User.Get(context.Background(), user.ID)
	if err != nil {
		log.Printf("failed getting user: %v", err)
	}
	user = client.User.GetX(context.Background(), user.ID)
	fmt.Printf("after Create\t%v\t%s\n", user, "saved correctly")

	user.Type = nil
	fmt.Printf("before Update\t%v\t\t%s\n", user, "type is nil")

	user.Update().SaveX(context.Background())
	user = client.User.GetX(context.Background(), user.ID)
	fmt.Printf("after Update\t%v\t%s\n", user, "nil value is not saved by Update alone")

	user.Update().ClearType().SaveX(context.Background())
	user = client.User.GetX(context.Background(), user.ID)
	fmt.Printf("after ClearType\t%v\t\t%s\n", user, "nil value is saved by ClearType")
}
