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
	log.Printf("after create\t%v\t%s", user, "proving that it was saved correctly")

	user.Type = nil
	log.Printf("before update\t%v\t\t%s", user, "note that type is nil")

	user.Update().SaveX(context.Background())
	user = client.User.GetX(context.Background(), user.ID)
	log.Printf("after update\t%v\t%s", user, "the nil value is not saved by (*ent.User).Update")

	user.Update().ClearType().SaveX(context.Background())
	user = client.User.GetX(context.Background(), user.ID)
	log.Printf("after cleartype\t%v\t\t%s", user, "now type is actually nil")
}
