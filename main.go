package main

import (
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/cyruzin/trailer/trailer/cmd"
)

func main() {
	client, err := tmdb.Init(os.Getenv("APIKey"))
	if err != nil {
		fmt.Println(err)
		return
	}

	trailer := cmd.NewTMDBClient(client)

	if err := trailer.RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
