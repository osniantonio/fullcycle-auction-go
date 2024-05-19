package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/osniantonio/fullcycle-auction-go/configuration/database/mongodb"
)

// rodar o mongo db via docker
// docker container run -d -p 27017:27017 --name auctionsDB mongo

func main() {
	ctx := context.Background()

	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println(err)
		log.Fatal("Error trying to load env variables")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
