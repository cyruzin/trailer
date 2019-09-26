package main

import (
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/cyruzin/trailer/trailer/cmd"
)

func main() {
	client, err := tmdb.Init("9aca69849a23528a419aea463387945f")
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
