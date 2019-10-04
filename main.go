package main

import (
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/cyruzin/trailer/trailer/cmd"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	apiKey, _ := os.LookupEnv("TMDB_KEY")

	client, err := tmdb.Init(apiKey)
	if err != nil {
		log.Println(err)
		return
	}

	trailer := cmd.NewTMDBClient(client)

	if err := trailer.RootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
