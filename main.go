package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	cmd "github.com/cyruzin/trailer/cmd/trailer"
)

func main() {

	// Declare cli flags
	useTMDB := flag.Bool("tmdb", false, "Use TheMovieDB API")
	searchMovie := flag.Bool("movie", false, "Search for Movies")
	searchTV := flag.Bool("tv", false, "Search for TV Shows")
	apiKey := flag.String("apiKey", "", "API key for The Movie DB")
	help := flag.Bool("help", false, "Shows the Help")
	flag.Parse()

	if *help {
		fmt.Println("usage:")
		fmt.Println("trailer -tmdb -apiKey=<yourApiKey> -movie Batman")
		fmt.Println("Available Flags:")
		flag.PrintDefaults()
		return
	}

	log.Printf("Using with Flags: -tmdb=%t, -movie=%t, -tv=%t, -apiKey=%s", *useTMDB, *searchMovie, *searchTV, *apiKey)

	if *apiKey == "" {
		apiKey, ok := os.LookupEnv("TMDB_KEY")
		if !ok {
			log.Println("TMDB api key is not")
			return
		}
		if apiKey == "" {
			log.Println("TMDB api key is empty")
			return
		}
	}

	client, err := tmdb.Init(*apiKey)
	if err != nil {
		log.Println(err)
		return
	}

	tClient := cmd.NewTMDBClient(client)

	if err := tClient.RootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
