package config

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent"
   _ "github.com/lib/pq"

   
   
)


func NewEntClient() *ent.Client {
//Open a connection to the database
   Client, err := ent.Open("postgres","host=localhost port=5432 user=postgres dbname=postgresDB password=mysecretpassword sslmode=disable")
   if err != nil {
      log.Fatal(err)
   }

   fmt.Println("Connected to database successfully")
   defer Client.Close()
// AutoMigration with ENT
   if err := Client.Schema.Create(context.Background()); err != nil {
      log.Fatalf("failed creating schema resources: %v", err)
   }
   return Client
}