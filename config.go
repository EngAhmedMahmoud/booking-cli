package main

import (
	"log"

	"github.com/joho/godotenv"
)
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load environment variables")
	}
}