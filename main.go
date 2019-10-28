package main

import (
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	cmd "github.com/cyruzin/trailer/cmd/trailer"
	"github.com/joho/godotenv"
)

var apiKey string

// Checking if the TMDb key is set in the environment variables.
func init() {
	apiKey = os.Getenv("TMDB_KEY")
}

// Fallback initialization in case the environment variable is not set.
//
// Checking if the TMDb key is set in the .env file.
func init() {
	if apiKey == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("[Trailer] No environment variable or .env file config was found")
			log.Println("")
		}

		apiKey, _ = os.LookupEnv("TMDB_KEY")
	}
}

func main() {
	log.SetFlags(0)

	client, err := tmdb.Init(apiKey)
	if err != nil {
		log.Println("[Wrapper]", err)
		return
	}

	trailer := cmd.NewTMDBClient(client)

	if err := trailer.RootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
