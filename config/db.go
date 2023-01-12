package config

import (
	"context"
	"fmt"
	"log"

	"github.com/george4joseph/go-blog-backend/ent"
	_ "github.com/lib/pq"
)

var ClientConfig *ent.Client

func NewEntClient() {
	//Open a connection to the database
	Client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=blogdb password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database successfully")
	// AutoMigration with ENT
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ClientConfig = Client
}
