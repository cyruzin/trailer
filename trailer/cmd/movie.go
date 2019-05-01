package cmd

import (
	"fmt"
	"os"
	"strings"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/spf13/cobra"
)

func movieCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "movie [name of the movie]",
		Short: "Searches for a movie trailer",
		Long:  "Searches one or more trailers from a movie.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := tmdb.Init(os.Getenv("APIKey"))

			if err != nil {
				fmt.Println(err)
			}

			argsJoin := strings.Join(args, " ")

			search, err := client.GetSearchMovies(argsJoin, nil)

			if err != nil {
				fmt.Println(err)
			}

			if len(search.Results) <= 0 {
				fmt.Println("No results for:", argsJoin)
				return
			}

			trailers, err := client.GetMovieVideos(int(search.Results[0].ID), nil)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Results for:", argsJoin)
			fmt.Println("")

			for _, trailer := range trailers.Results {
				fmt.Println(trailer.Name)
				fmt.Println(youtubeURL + trailer.Key)
				fmt.Println("")
			}
		},
	}
}
