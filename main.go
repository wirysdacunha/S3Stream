package main

import (
	"S3Stream/services"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found or failed to load.")
	}
	log.Println("Starting S3Stream service...")
	services.StartConsumer()
}
