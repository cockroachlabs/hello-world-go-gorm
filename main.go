package main

import (
	"bufio"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to the "bank" database as the "maxroach" user.
	// Read in connection string
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Enter a connection string: ")
	scanner.Scan()
	connstring := os.ExpandEnv(scanner.Text())

	// Connect to the "bank" database
	_, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	log.Println("Hey! You successfully connected to your CockroachDB cluster.")

}
